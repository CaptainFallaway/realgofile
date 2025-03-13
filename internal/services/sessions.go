package services

import (
	"sync"
	"time"

	"github.com/CaptainFallaway/realgofile/pkg/logging"
	"github.com/google/uuid"
)

// TODO: Idea: Instead of the implementation i have how about we create a object for each user so that I can bind tokens to the user directly and not only that see if theres state that one of the tokens are being used for a file upload.

const sessionDuration = time.Duration(10 * time.Minute)

type Session struct {
	Token     string    `json:"token"`
	UserId    string    `json:"uid"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
}

type SessionService struct {
	logger logging.Logger

	mux      sync.Mutex
	sessions map[string]*Session // I should take a deeper look into the sync.Map implementation
}

func NewSessionService(logger logging.Logger) *SessionService {
	return &SessionService{
		logger:   logger,
		sessions: make(map[string]*Session),
	}
}

func (ss *SessionService) newSession(uid string) (string, error) {
	token, err := uuid.NewV7()
	if err != nil {
		return "", nil
	}

	now := time.Now()

	session := &Session{
		Token:     token.String(),
		UserId:    uid,
		Expiry:    now.Add(sessionDuration),
		CreatedAt: now,
	}

	ss.mux.Lock()
	defer ss.mux.Unlock()

	ss.sessions[token.String()] = session

	return token.String(), nil
}

// GetSessionToken will return a new session token or a already existing session token for the user.
func (ss *SessionService) GetSessionToken(uid string) (string, error) {
	session := ss.GetSession(uid)
	if session == nil || !ss.IsValidSession(session.Token) {
		return ss.newSession(uid)
	}

	return session.Token, nil
}

func (ss *SessionService) IsValidSession(token string) bool {
	ss.mux.Lock()
	defer ss.mux.Unlock()

	session, ok := ss.sessions[token]

	if !ok {
		return false
	}

	if session.Expiry.Before(time.Now()) {
		delete(ss.sessions, token)
		return false
	}

	return true
}

// GetSession fetches the users session from the active sessions.
// This method will return nil if the session is no longer active.
func (ss *SessionService) GetSession(uid string) *Session {
	ss.mux.Lock()
	defer ss.mux.Unlock()

	for _, session := range ss.sessions {
		if session.UserId == uid {
			return session
		}
	}

	return nil
}
