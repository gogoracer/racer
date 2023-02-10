package engine

import (
	"context"
	"sync"
	"time"

	"github.com/gogoracer/racer/pkg/gas"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
)

type PageSession struct {
	Send             chan *gas.ToClient   // Buffered channel of outbound messages.
	Receive          chan *gas.FromClient // Buffered channel of inbound messages.
	id               uint64
	connected        bool
	connectedAt      time.Time
	lastActive       time.Time
	page             Pager
	ctxInitial       context.Context //nolint:containedctx // we are a router and create new contexts from this one
	ctxPage          context.Context //nolint:containedctx // we are a router and create new contexts from this one
	ctxPageCancel    context.CancelFunc
	ctxInitialCancel context.CancelFunc
	done             chan bool
	websocketConn    *websocket.Conn
	logger           zerolog.Logger
	muSess           sync.RWMutex
	muWrite          sync.RWMutex
	muRead           sync.RWMutex
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (sess *PageSession) readPump() {
	defer func() {
		sess.muSess.Lock()
		sess.connected = false
		sess.muSess.Unlock()

		sess.muWrite.Lock()
		if err := sess.websocketConn.Close(); err != nil {
			sess.logger.Err(err).Uint64("sess", sess.id).Msg("ws conn close")
		} else {
			sess.logger.Debug().Uint64("sess", sess.id).Msg("ws conn close")
		}
		sess.muWrite.Unlock()
	}()

	sess.muWrite.Lock()

	// c.conn.SetReadLimit(maxMessageSize)
	if err := sess.websocketConn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		sess.logger.Err(err).Msg("read pump set read deadline")
	}

	sess.websocketConn.SetPongHandler(func(string) error {
		sess.logger.Trace().Msg("ws pong")

		sess.muWrite.Lock()

		if err := sess.websocketConn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			sess.logger.Err(err).Msg("pong handler: set read deadline")
		}

		sess.muWrite.Unlock()

		return nil
	})

	sess.muWrite.Unlock()

	for {
		select {
		case <-sess.GetContextInitial().Done():
			return
		default:
			sess.muRead.Lock()

			mt, fromClientBytes, err := sess.websocketConn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					sess.logger.Debug().Err(err).Msg("unexpected close error")
				}
				return
			}

			if mt != websocket.BinaryMessage {
				sess.logger.Debug().Msg("not binary message")
				sess.websocketConn.Close()
			}

			sess.muRead.Unlock()

			sess.muSess.Lock()
			sess.lastActive = time.Now()
			sess.muSess.Unlock()

			fromClient := &gas.FromClient{}
			if err := fromClient.UnmarshalVT(fromClientBytes); err != nil {
				sess.logger.Err(err).Msg("read pump: unmarshal from client")
				return
			}

			sess.Receive <- fromClient
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (sess *PageSession) writePump() {
	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case <-sess.GetContextInitial().Done():
			ticker.Stop()

			return
		case toClient, ok := <-sess.Send:
			sess.muWrite.Lock()

			if err := sess.websocketConn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				sess.logger.Err(err).Msg("write pump: message set write deadline")
			}

			if !ok {
				// Send channel closed.
				if err := sess.websocketConn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					sess.logger.Err(err).Msg("write pump: write close message")
				}

				sess.muWrite.Unlock()

				return
			}

			w, err := sess.websocketConn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				sess.logger.Err(err).Msg("write pump: create writer")

				sess.muWrite.Unlock()

				continue
			}

			b, err := toClient.MarshalVT()
			if err != nil {
				sess.logger.Err(err).Msg("write pump: marshal to client")
			}

			if _, err := w.Write(b); err != nil {
				sess.logger.Err(err).Msg("write pump: write first message")
			}

			if err := w.Close(); err != nil {
				sess.logger.Err(err).Msg("write pump: close write")
			}

			sess.muWrite.Unlock()

		case <-ticker.C:
			sess.logger.Trace().Msg("ws ping")

			sess.muWrite.Lock()

			if err := sess.websocketConn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				sess.logger.Err(err).Msg("write pump: ping tick: set write deadline")
			}

			if err := sess.websocketConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				sess.logger.Err(err).Msg("write pump: ping tick: write write message")
			}

			sess.muWrite.Unlock()
		}
	}
}

func (sess *PageSession) GetPage() Pager {
	sess.muSess.RLock()
	defer sess.muSess.RUnlock()

	return sess.page
}

func (sess *PageSession) SetPage(page Pager) {
	sess.muSess.Lock()
	sess.page = page
	sess.muSess.Unlock()
}

func (sess *PageSession) GetID() uint64 {
	sess.muSess.RLock()
	defer sess.muSess.RUnlock()

	return sess.id
}

func (sess *PageSession) GetContextInitial() context.Context {
	sess.muSess.RLock()
	defer sess.muSess.RUnlock()

	return sess.ctxInitial
}

func (sess *PageSession) GetContextPage() context.Context {
	sess.muSess.RLock()
	defer sess.muSess.RUnlock()

	return sess.ctxPage
}

func (sess *PageSession) GetPageContextCancel() context.CancelFunc {
	sess.muSess.RLock()
	defer sess.muSess.RUnlock()

	return sess.ctxPageCancel
}

func (sess *PageSession) GetInitialContextCancel() context.CancelFunc {
	sess.muSess.RLock()
	defer sess.muSess.RUnlock()

	return sess.ctxInitialCancel
}

func (sess *PageSession) SetContextPage(ctx context.Context) {
	sess.muSess.Lock()
	sess.ctxPage = ctx
	sess.muSess.Unlock()
}

func (sess *PageSession) SetContextCancel(cancel context.CancelFunc) {
	sess.muSess.Lock()
	sess.ctxPageCancel = cancel
	sess.muSess.Unlock()
}

func (sess *PageSession) IsConnected() bool {
	sess.muSess.RLock()
	defer sess.muSess.RUnlock()

	return sess.connected
}

func (sess *PageSession) SetConnected(connected bool) {
	sess.muSess.Lock()
	sess.connected = connected
	sess.muSess.Unlock()
}
