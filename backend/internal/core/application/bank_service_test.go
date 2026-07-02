package application

import (
	"budging/backend/internal/core/domain"
	"testing"
)

type mockBankClient struct {
	aspsp []domain.Aspsp
	err   error
}

func (mockClient mockBankClient) GetAspsp() ([]domain.Aspsp, error) {
	return mockClient.aspsp, mockClient.err
}
func TestBankService_GetAspsp(t *testing.T) {
	bankClient := mockBankClient{
		aspsp: []domain.Aspsp{
			{Name: "bank-1", Country: "country-1"},
			{Name: "bank-2", Country: "country-2"},
		},
	}

	bankService := NewBankService(bankClient)
	result, err := bankService.GetAspsp()
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 aspsps not %d", len(result))
	}
	if result[0].Name != "bank-1" {
		t.Errorf("expected bank-1, got %s", result[0].Name)
	}
}
