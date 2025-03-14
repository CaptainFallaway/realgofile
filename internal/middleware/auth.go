package middleware

import "github.com/CaptainFallaway/realgofile/internal/services"

type AuthMiddleware struct {
	ss *services.SessionService
}

func NewAuthMiddleware(sessions *services.SessionService) *AuthMiddleware {
	return &AuthMiddleware{sessions}
}
