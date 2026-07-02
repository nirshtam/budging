package httpserver

import (
	"net/http"
)

func StartHttpServer(dependencies *Dependencies) error {
	mux := http.NewServeMux()
	bankingHandler := &BankingHandler{bankService: dependencies.BankService}
	RegisterBankingRoutes(mux, bankingHandler)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		return err
	}
	return nil
}
