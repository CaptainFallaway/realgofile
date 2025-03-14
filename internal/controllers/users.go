package controllers

import (
	"fmt"
	"net/http"

	"github.com/CaptainFallaway/realgofile/internal/services"
	"github.com/CaptainFallaway/realgofile/pkg/logging"
	"github.com/go-chi/chi/v5"
)

type userController struct {
	logger logging.Logger

	authService    *services.AuthService
	sessionService *services.SessionService
}

func NewUsers(logger logging.Logger, auth *services.AuthService, sessions *services.SessionService) Controller {
	return &userController{logger, auth, sessions}
}

func (u *userController) SetupRoutes(router chi.Router) {
	ld := NewLoggingDecorator(u.logger)

	router.Post("/login", ld.Decorate(u.Login))
	router.Post("/register", ld.Decorate(u.Register))
	router.Get("/sessions", ld.Decorate(u.GetSessions))
}

func (u *userController) Login(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		WriteError(w, "error parsing form", http.StatusBadRequest)
		return fmt.Errorf("error parsing form: %s", err)
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	verified, err := u.authService.Login(r.Context(), r.RemoteAddr, username, password)
	if err != nil {
		WriteError(w, "not authorized", http.StatusForbidden)
		return err
	}

	u.logger.Debug("successful login", "user", username, "ip", r.RemoteAddr)

	_, err = fmt.Fprint(w, verified)

	return err
}

func (u *userController) Register(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		WriteError(w, "error parsing form", http.StatusBadRequest)
		return fmt.Errorf("error parsing form: %s", err)
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	err = u.authService.Register(r.Context(), username, password)
	if err != nil {
		WriteError(w, "unable to register", http.StatusInternalServerError)
	}

	u.logger.Debug("successful registration", "user", username, "ip", r.RemoteAddr)

	return err
}

func (u *userController) GetSessions(w http.ResponseWriter, r *http.Request) error {
	sessions := u.sessionService.GetSessions()

	err := WriteJson(w, sessions)

	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
