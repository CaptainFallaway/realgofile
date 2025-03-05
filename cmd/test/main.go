package main

import (
	"context"
	"log"

	"github.com/CaptainFallaway/realgofile/internal/database"
	"github.com/CaptainFallaway/realgofile/pkg/hashing"
)

func main() {
	storage, err := database.NewSqliteRepo("data/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	h := hashing.NewArgonBcryptHasher()

	user, err := storage.GetUserByUsername(context.TODO(), "capnroot")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user.Username)
	log.Println(h.Compare("password", user.Salt, user.Password))
}
