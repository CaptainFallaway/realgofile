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

	authService *services.AuthService
}

func NewUsers(logger logging.Logger, auth *services.AuthService) Controller {
	return &userController{logger, auth}
}

func (u *userController) SetupRoutes(router chi.Router) {
	ld := NewLoggingDecorator(u.logger)

	router.Post("/login", ld.Decorate(u.Login))
	router.Post("/register", ld.Decorate(u.Register))
}

func (u *userController) Login(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		WriteError(w, "error parsing form", http.StatusBadRequest)
		return fmt.Errorf("error parsing form: %s", err)
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	u.logger.Info("login", "user", username, "pass", password)

	verified, err := u.authService.Login(r.Context(), r.RemoteAddr, username, password)
	if err != nil {
		WriteError(w, "not authorized", http.StatusForbidden)
		return err
	}

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

	u.logger.Info("register", "user", username, "pass", password, "ip", r.RemoteAddr)

	err = u.authService.Register(r.Context(), username, password)
	if err != nil {
		WriteError(w, "unable to register", http.StatusInternalServerError)
	}

	return err
}
