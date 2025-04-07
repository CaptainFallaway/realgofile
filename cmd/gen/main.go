package main

import (
	"os"

	"github.com/CaptainFallaway/realgofile/internal/storage"
	"github.com/CaptainFallaway/realgofile/pkg/hashing"
	"github.com/charmbracelet/log"
)

func main() {
	path, exist := os.LookupEnv("DBSTRING")
	if !exist {
		log.Fatal("dbstring not in env")
	}

	repo, err := storage.NewSqliteRepo(path)
	if err != nil {
		log.Fatal(err)
	}

	hasher := hashing.NewDefaultArgonBcryptHasher()

	app := NewAbstraction(repo, hasher)

	users, err := app.ListAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		log.Print(user.Username)
	}
}
