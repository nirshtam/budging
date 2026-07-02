package httpserver

import "net/http"

func RegisterBankingRoutes(mux *http.ServeMux, handler *BankingHandler) {
	mux.HandleFunc("GET /aspsps", handler.GetAspsp)
}
