package bank_client

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenProvider struct {
	AppID   string
	SignKey *rsa.PrivateKey
}

func (tokenProvider *TokenProvider) GetToken() (string, error) {
	if tokenProvider.AppID == "" {
		return "", fmt.Errorf("APP_ID is empty")
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"iss": "enablebanking.com",
		"aud": "api.enablebanking.com",
		"iat": now.Unix(),
		"exp": now.Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token.Header["typ"] = "JWT"
	token.Header["kid"] = tokenProvider.AppID

	signedToken, err := token.SignedString(tokenProvider.SignKey)
	return signedToken, err
}
