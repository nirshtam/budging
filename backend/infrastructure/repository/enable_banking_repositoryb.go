// Package repository
package repository

import (
	"budging/backend/infrastructure"
	"crypto/rsa"
	"io"
	"log"
	"net/http"
)

type EnableBankingRepository struct {
	baseURL string
	signKey *rsa.PrivateKey
}

func (repository EnableBankingRepository) GetAspsps(tokenProvider *infrastructure.TokenProvider) (*http.Response, error) {
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

func (repository EnableBankingRepository) Authenticate(tokenProvider *infrastructure.TokenProvider) {
	token, err := tokenProvider.GetToken()
	if err != nil {
		log.Fatal("Failed to issue token: ", err)
	}

	// reqBody := map[string]any{
	// 	"access": map[string]any{
	// 		"valid_until": time.Now().Add(90 * 24 * time.Hour).UTC().Format(time.RFC3339),
	// 	},
	// 	"aspsp": map[string]any{
	// 		"name":    "Fineco",
	// 		"country": "IT",
	// 	},
	// 	"state":        os.Getenv("appID"),
	// 	"redirect_url": redirectURL,
	// 	"psu_type":     "personal",
	// }

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
	if err != nil {
		log.Fatal("Error on response", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading the response bytes", err)
	}
	log.Println(string([]byte(body)))
}

func NewEnableBankingRepository(baseURL string, signKey *rsa.PrivateKey) *EnableBankingRepository {
	return &EnableBankingRepository{baseURL: baseURL, signKey: signKey}
}
