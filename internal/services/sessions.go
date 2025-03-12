package services

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/CaptainFallaway/realgofile/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
)

const (
	expiryDuration = time.Duration(30 * time.Minute)
	rsaKeySize     = 2048
)

type SessionService struct {
	key    *rsa.PrivateKey
	issuer string

	logger logging.Logger
}

// NewSessionService might panic if the generation of the rsa private key is not successful
func NewSessionService(logger logging.Logger) *SessionService {
	key, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		panic(err)
	}

	issuer := fmt.Sprintf("realgofile-%x", time.Now().Unix())

	return &SessionService{key, issuer, logger}
}

type jwtClaims struct {
	Uid string `json:"uid"`
	Ip  string `json:"ip"`
	jwt.RegisteredClaims
}

func (ss *SessionService) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected jwt signing method: %v", token.Header["alg"])
	}

	return ss.key, nil
}

// CreateSession returns a jwt that should be used for cookie based auth.
func (ss *SessionService) NewSession(uid, ip string) (string, error) {
	expiration := jwt.NewNumericDate(time.Now().Add(expiryDuration))

	claims := &jwtClaims{
		uid,
		ip,
		jwt.RegisteredClaims{
			ExpiresAt: expiration,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token, err := t.SignedString(ss.key)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (ss *SessionService) ProcessToken(sessionToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(sessionToken, &jwtClaims{}, ss.keyFunc, jwt.WithIssuer(ss.issuer), jwt.WithExpirationRequired())
	if err != nil {
		ss.logger.Error("jwt parse", "err", err)
		return false, ""
	}

	expiry, err := token.Claims.GetExpirationTime()
	if err != nil {
		ss.logger.Error("jwt get expiry", "err", err)
		return false, ""
	}

	claims, ok := token.Claims.(jwtClaims)
	if !ok {
		return false, ""
	}

	return expiry.Before(time.Now()) && token.Valid, claims.Uid
}
