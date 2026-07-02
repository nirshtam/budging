package application

import (
	"budging/backend/internal/core/domain"
	"budging/backend/internal/core/ports"
)

type BankService struct {
	bankClient ports.BankClient
}

func (service *BankService) GetAspsp() ([]domain.Aspsp, error) {
	aspspList, err := service.bankClient.GetAspsp()
	if err != nil {
		return nil, err
	}
	return aspspList, nil
}

func NewBankService(bankClient ports.BankClient) *BankService {
	return &BankService{bankClient: bankClient}
}
