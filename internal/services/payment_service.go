package services

import (
	"errors"

	"github.com/spacdust/merchant-bank-api/internal/models"
	"github.com/spacdust/merchant-bank-api/internal/repository"
)

type PaymentService struct {
	CustomerRepo *repository.CustomerRepository
	HistoryRepo  *repository.HistoryRepository
}

func NewPaymentService(cr *repository.CustomerRepository, hr *repository.HistoryRepository) *PaymentService {
	return &PaymentService{
		CustomerRepo: cr,
		HistoryRepo:  hr,
	}
}

func (s *PaymentService) MakePayment(fromID, toID string, amount float64) error {
	fromCustomer, err := s.CustomerRepo.GetByID(fromID)
	if err != nil {
		return err
	}
	if fromCustomer == nil {
		return errors.New("sender does not exist")
	}

	toCustomer, err := s.CustomerRepo.GetByID(toID)
	if err != nil {
		return err
	}
	if toCustomer == nil {
		return errors.New("recipient does not exist")
	}

	// cek saldo mencukupi
	if fromCustomer.Balance < amount {
		return errors.New("insufficient balance")
	}

	// tidak ada batasan transfer
	fromCustomer.Balance -= amount
	toCustomer.Balance += amount

	// update customer
	if err := s.CustomerRepo.Update(fromCustomer); err != nil {
		return err
	}
	if err := s.CustomerRepo.Update(toCustomer); err != nil {
		return err
	}

	// tambahkan ke history
	history := &models.History{
		ID:         generateID(),
		CustomerID: fromID,
		Action:     "payment",
		Amount:     amount,
		Timestamp:  getCurrentTime(),
	}

	return s.HistoryRepo.Add(history)
}
