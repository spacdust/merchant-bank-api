package services

import (
	"errors"
	"log"

	"github.com/spacdust/merchant-bank-api/internal/models"
	"github.com/spacdust/merchant-bank-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	CustomerRepo *repository.CustomerRepository
}

func NewAuthService(cr *repository.CustomerRepository) *AuthService {
	return &AuthService{CustomerRepo: cr}
}

func (s *AuthService) Login(id, password string) (*models.Customer, error) {
	customer, err := s.CustomerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		log.Printf("Login failed: customer with ID %s does not exist", id)
		return nil, errors.New("customer does not exist")
	}

	log.Printf("Attempting to login customer ID: %s", id)
	// bandingkan hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(customer.HashedPassword), []byte(password)); err != nil {
		log.Printf("Login failed: invalid credentials for customer ID %s", id)
		return nil, errors.New("invalid credentials")
	}
	log.Printf("Login successful for customer ID: %s", id)
	return customer, nil
}

func (s *AuthService) Logout() error {
	return nil
}

func (s *AuthService) CreateCustomer(id, name, password string, balance float64) (*models.Customer, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	customer := &models.Customer{
		ID:             id,
		Name:           name,
		HashedPassword: string(hashedPassword),
		Balance:        balance,
	}
	if err := s.CustomerRepo.AddCustomer(customer); err != nil {
		return nil, err
	}
	return customer, nil
}
