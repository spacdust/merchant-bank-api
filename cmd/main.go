package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spacdust/merchant-bank-api/internal/handlers"
	"github.com/spacdust/merchant-bank-api/internal/repository"
	"github.com/spacdust/merchant-bank-api/internal/services"
)

func main() {
	// inisialisasi repo
	customerRepo, err := repository.NewCustomerRepository("data/customers.json")
	if err != nil {
		log.Fatalf("Gagal memuat Customer Repository: %v", err)
	}

	historyRepo, err := repository.NewHistoryRepository("data/history.json")
	if err != nil {
		log.Fatalf("Gagal memuat History Repository: %v", err)
	}

	// inisialisasi services
	authService := services.NewAuthService(customerRepo)
	paymentService := services.NewPaymentService(customerRepo, historyRepo)

	// inisialisasi handlers
	apiHandler := handlers.NewAPIHandler(authService, paymentService)

	// setup router
	router := mux.NewRouter()
	handlers.RegisterRoutes(router, apiHandler)

	// jalankan server
	log.Println("Server berjalan di :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server gagal dijalankan: %v", err)
	}
}
