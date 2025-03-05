package services

import "context"

type TotpService interface {
	TotpCreator
	TotpMan
}

type TotpCreator interface {
	Create(ctx context.Context, username, password string) (string, error)
}

type TotpMan interface {
	Verify(ctx context.Context, username, key string) (bool, error)
}

type totpService struct {
	issuer string
	crypto CryptService
}

func NewTotpService(issuer string, crypto CryptService) TotpService {
	return &totpService{issuer, crypto}
}

func (s *totpService) Create(ctx context.Context, username, key string) (string, error) {
	return "", nil
}

func (s *totpService) Verify(ctx context.Context, username, key string) (bool, error) {
	return false, nil
}

func (s *totpService) TotpEnabled(ctx context.Context, username string) (bool, error) {
	return false, nil
}
