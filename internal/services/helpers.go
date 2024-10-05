package services

import (
	"time"

	"github.com/google/uuid"
)

func generateID() string {
	return uuid.New().String()
}

func getCurrentTime() time.Time {
	return time.Now()
}
