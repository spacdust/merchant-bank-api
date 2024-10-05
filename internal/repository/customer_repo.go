package repository

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/spacdust/merchant-bank-api/internal/models"
)

type CustomerRepository struct {
	FilePath  string
	Mutex     sync.Mutex
	Customers []models.Customer
}

func NewCustomerRepository(filePath string) (*CustomerRepository, error) {
	repo := &CustomerRepository{FilePath: filePath}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &repo.Customers); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *CustomerRepository) GetByID(id string) (*models.Customer, error) {
	for _, customer := range r.Customers {
		if customer.ID == id {
			c := customer
			return &c, nil
		}
	}
	return nil, nil
}

func (r *CustomerRepository) Update(customer *models.Customer) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	for i, c := range r.Customers {
		if c.ID == customer.ID {
			r.Customers[i] = *customer
			break
		}
	}
	data, err := json.MarshalIndent(r.Customers, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.FilePath, data, 0644)
}

func (r *CustomerRepository) AddCustomer(customer *models.Customer) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Customers = append(r.Customers, *customer)
	data, err := json.MarshalIndent(r.Customers, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.FilePath, data, 0644)
}
