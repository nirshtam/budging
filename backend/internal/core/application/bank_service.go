package application

import (
	"budging/backend/internal/core/domain"
	"budging/backend/internal/core/ports"
)

type BankService struct {
	BankClient ports.BankClient
}

func (service *BankService) GetAspsp() ([]domain.Aspsp, error) {
	aspspList, err := service.BankClient.GetAspsp()
	if err != nil {
		return nil, err
	}
	return aspspList, nil
}

func NewBankService(bankClient ports.BankClient) *BankService {
	return &BankService{BankClient: bankClient}
}
