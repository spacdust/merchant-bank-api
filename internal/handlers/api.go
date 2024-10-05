package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spacdust/merchant-bank-api/internal/services"
)

type APIHandler struct {
	AuthService    *services.AuthService
	PaymentService *services.PaymentService
}

func NewAPIHandler(as *services.AuthService, ps *services.PaymentService) *APIHandler {
	return &APIHandler{
		AuthService:    as,
		PaymentService: ps,
	}
}

// handler untuk Login
func (h *APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	customer, err := h.AuthService.Login(creds.ID, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]string{"message": "login successful", "customer_id": customer.ID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handler payment
func (h *APIHandler) Payment(w http.ResponseWriter, r *http.Request) {
	var payment struct {
		FromID string  `json:"from_id"`
		ToID   string  `json:"to_id"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if payment.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}

	if err := h.PaymentService.MakePayment(payment.FromID, payment.ToID, payment.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{"message": "payment successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handler logout
func (h *APIHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Panggil metode Logout dari AuthService
	if err := h.AuthService.Logout(); err != nil {
		http.Error(w, "Logout failed", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "logout successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// daftarkan semua rute
func RegisterRoutes(router *mux.Router, handler *APIHandler) {
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/payment", handler.Payment).Methods("POST")
	router.HandleFunc("/logout", handler.Logout).Methods("POST")
}
