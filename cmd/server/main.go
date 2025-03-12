package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CaptainFallaway/realgofile/internal/controllers"
	"github.com/CaptainFallaway/realgofile/internal/services"
	"github.com/CaptainFallaway/realgofile/internal/storage"
	"github.com/CaptainFallaway/realgofile/pkg/hashing"
	"github.com/CaptainFallaway/realgofile/pkg/logging"
	"github.com/go-chi/chi/v5"
)

const addr = "0.0.0.0:3000"

func main() {
	logger := logging.NewCharmLogger(os.Stdout)

	path, exist := os.LookupEnv("DBSTRING")
	if !exist {
		log.Fatal("db path not in env")
	}

	repo, err := storage.NewSqliteRepo(path)
	if err != nil {
		logger.Fatal(err)
	}

	hasher := hashing.NewArgonBcryptHasher()

	ss := services.NewSessionService(logger)

	authService := services.NewAuthService(repo, hasher, ss)

	userController := controllers.NewUsers(logger, authService)

	r := chi.NewRouter()

	r.Route("/user", userController.SetupRoutes)

	logger.Info("server starting", "addr", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		logger.Fatal(err)
	}
}
