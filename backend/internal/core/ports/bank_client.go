package ports

import (
	"budging/backend/internal/core/domain"
)

type BankClient interface {
	GetAspsp() ([]domain.Aspsp, error)
}
