package tests

import (
	"testing"

	"github.com/spacdust/merchant-bank-api/internal/models"
	"github.com/spacdust/merchant-bank-api/internal/repository"
	"github.com/spacdust/merchant-bank-api/internal/services"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	// Setup
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	customers := []models.Customer{
		{
			ID:             "cust1",
			Name:           "Alice",
			HashedPassword: string(hashedPassword),
			Balance:        1000.0,
		},
	}
	customerRepo := &repository.CustomerRepository{
		FilePath:  "data/customers_test.json",
		Customers: customers,
	}
	authService := services.NewAuthService(customerRepo)

	// Test successful login
	customer, err := authService.Login("cust1", "password123")
	if err != nil || customer.ID != "cust1" {
		t.Errorf("Expected successful login, got error: %v", err)
	}

	// Test invalid password
	_, err = authService.Login("cust1", "wrongpass")
	if err == nil {
		t.Errorf("Expected login failure, got success")
	}

	// Test non-existent customer
	_, err = authService.Login("cust2", "password123")
	if err == nil {
		t.Errorf("Expected login failure for non-existent customer, got success")
	}
}
