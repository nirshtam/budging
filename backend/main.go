package main

import (
	"budging/backend/internal/adapters/enablebanking"
	"budging/backend/internal/adapters/httpserver"
	"budging/backend/internal/core/application"
	"crypto/rsa"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	var signKey *rsa.PrivateKey

	keyPath := os.Getenv("PEM_KEY_PATH")
	key, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatal("Failed to read sign key: ", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Fatal("Invalid sign key")
	}

	enablebankingTokenProvider := &enablebanking.TokenProvider{
		AppID:   os.Getenv("APP_ID"),
		SignKey: signKey,
	}
	bankClient := &enablebanking.EnableBankingClient{
		BaseURL:       os.Getenv("ENABLEBANKING_BASE_URL"),
		TokenProvider: enablebankingTokenProvider,
	}
	dependencies := httpserver.NewDependencies(
		application.NewBankService(bankClient),
	)

	httpserver.StartHttpServer(dependencies)
}
