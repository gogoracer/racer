package engine

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/vmihailenco/msgpack/v5"
)

type PageFunc func() *Page

func NewPageServer(pf func() *Page) (*PageServer, error) {
	store, err := NewPageSessionStore()
	if err != nil {
		return nil, fmt.Errorf("failed to create page session store: %w", err)
	}
	p := NewPageServerWithSessionStore(pf, store)
	return p, nil
}

func MustNewPageServer(pf PageFunc) *PageServer {
	p, err := NewPageServer(pf)
	if err != nil {
		panic(err)
	}
	return p
}

func NewPageServerWithSessionStore(pf PageFunc, sess *PageSessionStore) *PageServer {
	return &PageServer{
		pageFunc: pf,
		Sessions: sess,
		logger:   zerolog.Nop(),
	}
}

type PageServer struct {
	Sessions *PageSessionStore
	Upgrader websocket.Upgrader

	pageFunc func() *Page
	logger   zerolog.Logger
}

func (s *PageServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// WebSocket?
	sessID := r.URL.Query().Get("hlive")

	if sessID == "" {
		s.pageFunc().ServeHTTP(w, r)

		return
	}

	var (
		sess *PageSession
		err  error
	)

	// New
	if sessID == "1" {
		sess, err = s.Sessions.New()
		if err != nil {
			s.logger.Error().Err(err).Msg("ws connect: failed to create new session")
			return
		}
		sess.muSess.Lock()
		sess.page = s.pageFunc()
		sess.connectedAt = time.Now()
		sess.lastActive = sess.connectedAt
		sess.ctxInitial, sess.ctxInitialCancel = context.WithCancel(r.Context())
		sess.ctxPage, sess.ctxPageCancel = context.WithCancel(sess.ctxInitial)
		sess.muSess.Unlock()
	} else { // Reconnect
		_ = "todo"
		// TODO: need to rethink reconnect and double check my assumptions
		//sess = s.Sessions.Get(sessID)
		//
		//if sess != nil && sess.IsConnected() {
		//	LoggerDev.Error().Str("sessID", sessID).
		//		Msg("ws connect: is connected: connection blocked as an active connection exists")
		//
		//	w.WriteHeader(http.StatusNotFound)
		//
		//	return
		//}

		//if sess != nil {
		//
		//	//sess.GetInitialContextCancel()()
		//
		//	//sess.muSess.Lock()
		//
		//	//sess.ctxInitial, sess.ctxInitialCancel = context.WithCancel(r.Context())
		//	//sess.ctxPage, sess.ctxPageCancel = context.WithCancel(sess.ctxInitial)
		//
		//	//sess.muSess.Unlock()
		//}
	}

	// LoggerDev.Debug().Str("sessID", sessID).Bool("success", sess != nil).Msg("ws connect")

	if sess == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	hhash := r.URL.Query().Get("hhash")

	s.logger = sess.GetPage().logger
	s.logger.Debug().Str("sessionID", sessID).Str("hash", hhash).Msg("ws start")

	if sess.GetPage().cache != nil && hhash != "" && sessID == "1" {
		val, hit := sess.GetPage().cache.Get(hhash)

		b, ok := val.([]byte)
		if hit && ok {
			s.logger.Debug().Bool("hit", hit).Str("hhash", hhash).Int("size", len(b)/1024).
				Msg("cache get")
			newTree := G()
			if err := msgpack.Unmarshal(b, newTree); err != nil {
				s.logger.Err(err).Msg("ServeHTTP: msgpack.Unmarshal")
			} else {
				sess.GetPage().domBrowser = newTree
			}
		}
	}

	sess.muSess.Lock()

	sess.websocketConn, err = s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		sess.muSess.Unlock()
		s.logger.Err(err).Msg("ws upgrade")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	sess.connected = true
	sess.lastActive = time.Now()

	sess.muSess.Unlock()

	go sess.writePump()
	go sess.readPump()

	if err := sess.GetPage().ServeWS(sess.GetContextPage(), sess.GetID(), sess.Send, sess.Receive); err != nil {
		sess.GetPage().logger.Err(err).Msg("ws serve")
		w.WriteHeader(http.StatusInternalServerError)
	}

	// This needs to say open to keep the context active
	<-sess.done
}
