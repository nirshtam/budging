package httpserver

import (
	"budging/backend/internal/core/application"
	"net/http"
)

type BankingHandler struct {
	bankService *application.BankService
}

func (handler *BankingHandler) GetAspsp(w http.ResponseWriter, r *http.Request) {
	handler.bankService.GetAspsp()
}
