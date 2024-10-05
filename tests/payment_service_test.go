package tests

import (
	"testing"

	"github.com/spacdust/merchant-bank-api/internal/models"
	"github.com/spacdust/merchant-bank-api/internal/repository"
	"github.com/spacdust/merchant-bank-api/internal/services"
)

func TestMakePayment(t *testing.T) {
	// setup
	customers := []models.Customer{
		{ID: "cust1", Name: "Alice", Balance: 1000.0},
		{ID: "cust2", Name: "Bob", Balance: 500.0},
	}
	customerRepo := &repository.CustomerRepository{
		FilePath:  "../data/customers_test.json",
		Customers: customers,
	}
	historyRepo := &repository.HistoryRepository{
		FilePath:  "../data/history_test.json",
		Histories: []models.History{},
	}
	paymentService := services.NewPaymentService(customerRepo, historyRepo)

	// test successful payment
	err := paymentService.MakePayment("cust1", "cust2", 200.0)
	if err != nil {
		t.Errorf("Expected successful payment, got error: %v", err)
	}
	updatedFrom, _ := customerRepo.GetByID("cust1")
	updatedTo, _ := customerRepo.GetByID("cust2")
	if updatedFrom.Balance != 800.0 || updatedTo.Balance != 700.0 {
		t.Errorf("Balances not updated correctly: From=%v, To=%v", updatedFrom.Balance, updatedTo.Balance)
	}

	// test payment with insufficient balance
	err = paymentService.MakePayment("cust1", "cust2", 1000.0)
	if err == nil {
		t.Errorf("Expected payment failure due to insufficient balance, got success")
	}

	// test payment to non-existent customer
	err = paymentService.MakePayment("cust1", "cust3", 100.0)
	if err == nil {
		t.Errorf("Expected payment failure to non-existent customer, got success")
	}
}
