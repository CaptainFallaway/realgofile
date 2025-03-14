package services

import (
	"sync"
	"time"

	"github.com/CaptainFallaway/realgofile/pkg/logging"
	"github.com/google/uuid"
)

// TODO: Idea: Instead of the implementation i have how about we create a object for each user so that I can bind tokens to the user directly and not only that see if theres state that one of the tokens are being used for a file upload.

const idleDuration = time.Duration(30 * time.Minute)

type Session struct {
	Uid        string    `json:"uid"`
	LastActive time.Time `json:"last_active"`
}

type SessionService struct {
	logger logging.Logger

	mux         sync.Mutex
	activeUsers map[string]string
	sessions    map[string]*Session
}

func NewSessionService(logger logging.Logger) *SessionService {
	return &SessionService{
		logger:      logger,
		activeUsers: make(map[string]string),
		sessions:    make(map[string]*Session),
	}
}

func (ss *SessionService) Login(uid string) (string, error) {
	token, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	session := &Session{
		Uid:        uid,
		LastActive: time.Now(),
	}

	ss.mux.Lock()
	defer ss.mux.Unlock()

	if otherToken, ok := ss.activeUsers[uid]; ok {
		delete(ss.sessions, otherToken)
	}

	ss.activeUsers[uid] = token.String()
	ss.sessions[token.String()] = session

	return token.String(), nil
}

func (ss *SessionService) Logout(token string) {
	ss.mux.Lock()
	defer ss.mux.Unlock()

	session, ok := ss.sessions[token]
	if ok {
		delete(ss.activeUsers, session.Uid)
	}

	delete(ss.sessions, token)
}

func (ss *SessionService) sessionIsExpired(session *Session) bool {
	return session.LastActive.Add(idleDuration).After(time.Now())
}

func (ss *SessionService) Authorize(token string) bool {
	ss.mux.Lock()
	defer ss.mux.Unlock()

	session, ok := ss.sessions[token]
	if !ok || ss.sessionIsExpired(session) {
		return false
	}

	session.LastActive = time.Now()

	return true
}

func (ss *SessionService) GetSessions() map[string]Session {
	ret := make(map[string]Session, len(ss.sessions))

	ss.mux.Lock()
	defer ss.mux.Unlock()

	for key, val := range ss.sessions {
		ret[key] = *val
	}

	return ret
}
