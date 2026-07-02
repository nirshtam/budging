package main

import (
	// "crypto/rsa"
	// "fmt"
	// "io"
	// "log"
	"net/http"
	// "os"
	// "budging/backend/infrastructure"
	// "budging/backend/infrastructure/repository"
	// "github.com/golang-jwt/jwt/v5"
)

func main() {
	// var signKey *rsa.PrivateKey
	// var err error
	// keyPath := os.Getenv("PEM_KEY_PATH")
	// key, err := os.ReadFile(keyPath)
	// if err != nil {
	// 	log.Fatal("Failed to read sign key: ", err)
	// }
	// signKey, err = jwt.ParseRSAPrivateKeyFromPEM(key)
	// if err != nil {
	// 	log.Fatal("Invalid sign key")
	// }

	// tokenProvider := infrastructure.TokenProvider{AppID: os.Getenv("appID"), SignKey: signKey}
	// repository := repository.NewEnableBankingRepository("https://api.enablebanking.com", signKey)

	http.ListenAndServe(":8000", nil)
	// resp, err := repository.GetAspsp(&tokenProvider)
	// if err != nil {
	// 	log.Fatal("Failed to retrieve aspsps")
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("Failed to read response body: ", err)
	// }
	// fmt.Println(string(body))
}
