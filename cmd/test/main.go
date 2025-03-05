package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CaptainFallaway/realgofile/internal/database"
	"github.com/CaptainFallaway/realgofile/pkg/hasher"
)

func main() {
	storage, err := database.NewSqliteRepo("data/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	h := hasher.NewArgonBcryptHasher()

	user, err := storage.GetUserByUsername(context.TODO(), "captainfallaway")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	fmt.Println(h.Compare("password", user.Salt, user.Password))
}
