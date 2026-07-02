package httpserver

import "budging/backend/internal/core/application"

type Dependencies struct {
	BankService *application.BankService
}

func NewDependencies(bankService *application.BankService) *Dependencies {
	return &Dependencies{BankService: bankService}
}
