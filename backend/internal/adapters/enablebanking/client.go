package enablebanking

import (
	"budging/backend/internal/core/domain"
	"crypto/rsa"
	"io"
	"log"
	"net/http"
)

type EnableBankingRepository struct {
	baseURL       string
	tokenProvider *TokenProvider
	mapper        *Mapper
}

func (repository EnableBankingRepository) GetAspsp() ([]domain.Aspsp, error) {
	token, err := repository.tokenProvider.GetToken()
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
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return repository.mapper.MapAspsp(string(body))
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

func NewEnableBankingClient(baseURL string, signKey *rsa.PrivateKey, appId string) *EnableBankingRepository {
	return &EnableBankingRepository{baseURL: baseURL, tokenProvider: &TokenProvider{AppID: appId, SignKey: signKey}}
}
