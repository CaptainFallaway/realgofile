package main

import (
	"context"

	"github.com/CaptainFallaway/realgofile/internal/storage"
	"github.com/CaptainFallaway/realgofile/pkg/hashing"
	"github.com/google/uuid"
)

const saltSize = 32

type Abstraction struct {
	ctx    context.Context
	repo   storage.Repository
	hasher hashing.HasherWithSalt
}

func NewAbstraction(repo storage.Repository, hasher hashing.HasherWithSalt) *Abstraction {
	ctx := context.Background()
	return &Abstraction{ctx, repo, hasher}
}

func (ga *Abstraction) NewUser(username, password string) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return nil
	}

	salt := hashing.GenerateSalt(saltSize)

	hashedPass, err := ga.hasher.Hash(password, salt)

	user := storage.User{
		Uid:      uid.String(),
		Username: username,
		Password: hashedPass,
		Salt:     salt,
	}

	err = ga.repo.InsertUser(ga.ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (ga *Abstraction) DeleteUser(username string) error {
	user, err := ga.repo.GetUserByUsername(ga.ctx, username)
	if err != nil {
		return err
	}

	err = ga.repo.DeleteUser(ga.ctx, user.Uid)
	return err
}

func (ga *Abstraction) ListAllUsers() ([]storage.User, error) {
	return ga.repo.GetAllUsers(ga.ctx)
}
