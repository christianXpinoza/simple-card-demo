package api

import (
	"net/http"

	"github.com/christianXpinoza/simple-card/card"
	"github.com/go-chi/chi"
)

// Service abstraction
type Service struct {
	router *chi.Mux
}

var cardService *card.Service

// Start the API service
func (s *Service) Start(addr string, cService *card.Service) error {

	// Card Service Initialization
	cardService = cService

	// Clients API portion
	s.router = chi.NewRouter()
	s.router.Post("/apiv1/clients/new", newClientHandler)
	s.router.Post("/apiv1/clients/deposit", depositHandler)
	s.router.Get("/apiv1/clients/balance", balanceHandler)

	// Merchants API portion
	s.router.Post("/apiv1/transaction/block_auth", blockAuthHandler)
	s.router.Delete("/apiv1/transaction/block_auth", deleteCaptureAuthHandler)
	s.router.Post("/apiv1/transaction/capture", captureHandler)

	return http.ListenAndServe(addr, s.router)
}

// New return a new instance of the API service
func New() *Service {
	a := new(Service)
	return a
}
