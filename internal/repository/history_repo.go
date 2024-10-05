package repository

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/spacdust/merchant-bank-api/internal/models"
)

type HistoryRepository struct {
	FilePath  string
	Mutex     sync.Mutex
	Histories []models.History
}

func NewHistoryRepository(filePath string) (*HistoryRepository, error) {
	repo := &HistoryRepository{FilePath: filePath}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &repo.Histories); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *HistoryRepository) Add(history *models.History) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Histories = append(r.Histories, *history)
	data, err := json.MarshalIndent(r.Histories, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.FilePath, data, 0644)
}
