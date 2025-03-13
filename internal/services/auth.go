package services

import (
	"context"
	"errors"

	"github.com/CaptainFallaway/realgofile/internal/storage"
	"github.com/CaptainFallaway/realgofile/pkg/hashing"
	"github.com/google/uuid"
)

type AuthService struct {
	repo   storage.Repository
	hasher hashing.HasherWithSalt

	sessions *SessionService
}

const saltSize = 32

var (
	ErrUnauthorized = errors.New("unauthorized")
)

func NewAuthService(repo storage.Repository, hasher hashing.HasherWithSalt, ss *SessionService) *AuthService {
	return &AuthService{repo, hasher, ss}
}

func (au *AuthService) Login(ctx context.Context, ip, username, password string) (string, error) {
	// TODO: Should probably have DTO for the data needed in this function
	user, err := au.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !au.hasher.Compare(password, user.Salt, user.Password) {
		return "", ErrUnauthorized
	}

	token, err := au.sessions.GetSessionToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (au *AuthService) Register(ctx context.Context, username, password string) error {
	salt := hashing.GenerateSalt(saltSize)

	hashedPass, err := au.hasher.Hash(password, salt)
	if err != nil {
		return err
	}

	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	user := storage.User{
		Uid:      uid.String(),
		Username: username,
		Password: hashedPass,
		Salt:     salt,
	}

	return au.repo.InsertUser(ctx, user)
}
