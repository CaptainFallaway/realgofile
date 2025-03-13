package main

import (
	"net/http"
	"os"

	"github.com/CaptainFallaway/realgofile/internal/config"
	"github.com/CaptainFallaway/realgofile/internal/controllers"
	"github.com/CaptainFallaway/realgofile/internal/services"
	"github.com/CaptainFallaway/realgofile/internal/storage"
	"github.com/CaptainFallaway/realgofile/pkg/hashing"
	"github.com/CaptainFallaway/realgofile/pkg/logging"
	"github.com/go-chi/chi/v5"
)

var conf = &config.Config{
	Addr:     ":3000",
	Debug:    true,
	DbString: "./data/db.sqlite3",
}

func main() {
	config.LoadEnv(conf)

	logger := logging.NewCharmLogger(os.Stderr, conf.Debug)

	repo, err := storage.NewSqliteRepo(conf.DbString)
	if err != nil {
		logger.Fatal(err)
	}

	hasher := hashing.NewArgonBcryptHasher()

	ss := services.NewSessionService(logger)

	authService := services.NewAuthService(repo, hasher, ss)

	userController := controllers.NewUsers(logger, authService)

	r := chi.NewRouter()

	r.Route("/user", userController.SetupRoutes)

	logger.Info("server starting", "addr", conf.Addr)
	err = http.ListenAndServe(conf.Addr, r)
	if err != nil {
		logger.Fatal(err)
	}
}
