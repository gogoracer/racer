package engine

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/cornelk/hashmap"
	"github.com/gogoracer/racer/pkg/gas"
	"github.com/rs/zerolog"
	"github.com/tpaschalis/daffodil"
)

type PageSessionStore struct {
	sessions              *hashmap.Map[uint64, *PageSession]
	DisconnectTimeout     time.Duration
	SessionLimit          uint32
	sessionCount          uint32
	GarbageCollectionTick time.Duration
	Done                  chan bool
	idGen                 func() uint64
}

func NewPageSessionStore() (*PageSessionStore, error) {
	idCfg, err := daffodil.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create session id generator config: %w", err)
	}
	idGen, err := daffodil.NewDaffodil(idCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create session id generator: %w", err)
	}

	pss := &PageSessionStore{
		DisconnectTimeout:     WebSocketDisconnectTimeoutDefault,
		SessionLimit:          PageSessionLimitDefault,
		GarbageCollectionTick: PageSessionGarbageCollectionTick,
		Done:                  make(chan bool),
		sessions:              hashmap.New[uint64, *PageSession](),
		idGen: func() uint64 {
			id, err := idGen.Next()
			if err != nil {
				panic(fmt.Errorf("failed to generate session id: %w", err))
			}
			return uint64(id)
		},
	}

	go pss.GarbageCollection()

	return pss, nil
}

// New PageSession.
func (pss *PageSessionStore) New() (*PageSession, error) {
	// Block until we have room for a new session
	pss.newWait()

	ps := &PageSession{
		id:      pss.idGen(),
		logger:  zerolog.Nop(),
		Send:    make(chan *gas.ToClient),
		Receive: make(chan *gas.FromClient),
		done:    make(chan bool),
	}

	pss.mapAdd(ps)

	return ps, nil
}

// TODO: use sync.Cond
func (pss *PageSessionStore) newWait() {
	for atomic.LoadUint32(&pss.sessionCount) > pss.SessionLimit {
		runtime.Gosched()
	}
}

func (pss *PageSessionStore) Get(id uint64) *PageSession {
	return pss.mapGet(id)
}

func (pss *PageSessionStore) mapAdd(ps *PageSession) {
	pss.sessions.Set(ps.id, ps)
	atomic.AddUint32(&pss.sessionCount, 1)
}

func (pss *PageSessionStore) mapGet(id uint64) *PageSession {
	ps, _ := pss.sessions.Get(id)

	return ps
}

func (pss *PageSessionStore) mapDelete(id uint64) {
	if pss.sessions.Del(id) {
		atomic.AddUint32(&pss.sessionCount, ^uint32(0))
	}
}

func (pss *PageSessionStore) GarbageCollection() {
	for {
		time.Sleep(pss.GarbageCollectionTick)

		select {
		case <-pss.Done:
			return
		default:
			now := time.Now()
			pss.sessions.Range(func(id uint64, sess *PageSession) bool {
				if sess.IsConnected() {
					return true
				}

				// Keep until it exceeds the timeout
				sess.muSess.RLock()
				la := sess.lastActive
				sess.muSess.RUnlock()

				if now.Sub(la) > pss.DisconnectTimeout {
					if sess.page != nil {
						sess.page.Close(sess.ctxPage)
					}

					if sess.ctxInitialCancel != nil {
						sess.ctxInitialCancel()
					}

					close(sess.done)
					pss.mapDelete(id)
				}

				return true
			})
		}
	}
}

func (pss *PageSessionStore) Delete(id uint64) {
	ps := pss.mapGet(id)
	if ps == nil {
		return
	}

	if ps.GetPage() != nil {
		ps.GetPage().Close(ps.GetContextPage())
	}

	pss.mapDelete(id)
}

func (pss *PageSessionStore) GetSessionCount() int {
	return int(atomic.LoadUint32(&pss.sessionCount))
}
