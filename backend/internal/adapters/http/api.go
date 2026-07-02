package http

import (
	"budging/backend/internal/core/application"
	"log"
	"net/http"
)

func StartHttpServer(bankService *application.BankService) {
	mux := http.NewServeMux()
	bankingHandler := &BankingHandler{bankService: bankService}
	RegisterBankingRoutes(mux, bankingHandler)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("Unable to start http server: ", err)
	}
}
