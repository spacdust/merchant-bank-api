package repository

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/spacdust/merchant-bank-api/internal/models"
)

type MerchantRepository struct {
	FilePath  string
	Mutex     sync.Mutex
	Merchants []models.Merchant
}

func NewMerchantRepository(filePath string) (*MerchantRepository, error) {
	repo := &MerchantRepository{FilePath: filePath}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &repo.Merchants); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *MerchantRepository) GetByID(id string) (*models.Merchant, error) {
	for _, merchant := range r.Merchants {
		if merchant.ID == id {
			m := merchant
			return &m, nil
		}
	}
	return nil, nil
}
