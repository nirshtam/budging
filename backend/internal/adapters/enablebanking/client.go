// Package bank_client
package bank_client

import (
	"crypto/rsa"
	"log"
	"net/http"
)

type EnableBankingRepository struct {
	baseURL string
	signKey *rsa.PrivateKey
}

func (repository EnableBankingRepository) GetAspsps(tokenProvider *TokenProvider) (*http.Response, error) {
	token, err := tokenProvider.GetToken()
	if err != nil {
		log.Fatal("Failed to issue token: ", err)
	}

	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", repository.baseURL+"/aspsps", nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Host", "api.enablebanking.com")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

func (repository EnableBankingRepository) Authenticate(tokenProvider *TokenProvider) (*http.Response, error) {
	token, err := tokenProvider.GetToken()
	if err != nil {
		log.Fatal("Failed to issue token: ", err)
	}

	bearer := "Bearer " + token
	req, err := http.NewRequest("POST", repository.baseURL+"/auth", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Host", "api.enablebanking.com")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

func NewEnableBankingRepository(baseURL string, signKey *rsa.PrivateKey) *EnableBankingRepository {
	return &EnableBankingRepository{baseURL: baseURL, signKey: signKey}
}
