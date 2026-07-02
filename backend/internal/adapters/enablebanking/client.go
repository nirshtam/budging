package enablebanking

import (
	"budging/backend/internal/core/domain"
	"crypto/rsa"
	"io"
	"log"
	"net/http"
)

type EnableBankingClient struct {
	BaseURL       string
	TokenProvider *TokenProvider
}

func (enablebankingClient *EnableBankingClient) GetAspsp() ([]domain.Aspsp, error) {
	token, err := enablebankingClient.TokenProvider.GetToken()
	if err != nil {
		log.Fatal("Failed to issue token: ", err)
	}

	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", enablebankingClient.BaseURL+"/aspsps", nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Host", "api.enablebanking.com")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return MapAspsp(string(body))
}

func (enablebankingClient *EnableBankingClient) Authenticate() (*http.Response, error) {
	token, err := enablebankingClient.TokenProvider.GetToken()
	if err != nil {
		log.Fatal("Failed to issue token: ", err)
	}

	bearer := "Bearer " + token
	req, err := http.NewRequest("POST", enablebankingClient.BaseURL+"/auth", nil)
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

func NewEnableBankingClient(baseURL string, signKey *rsa.PrivateKey, appId string) *EnableBankingClient {
	return &EnableBankingClient{BaseURL: baseURL, TokenProvider: &TokenProvider{AppID: appId, SignKey: signKey}}
}
