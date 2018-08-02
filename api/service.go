package api

import (
	"net/http"

	"github.com/christianXpinoza/simple-card-demo/card"
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

	s.router.Post("/apiv1/cards/new", newClientHandler)
	s.router.Get("/apiv1/cards/balance", balanceHandler)
	s.router.Get("/apiv1/cards/statement", transactionStatementHandler)

	s.router.Post("/apiv1/claim/block_auth", blockAuthHandler)
	s.router.Delete("/apiv1/claim/block_auth", cancelBlockingAuthHandler)

	s.router.Post("/apiv1/transaction/deposit", depositHandler)
	s.router.Post("/apiv1/transaction/capture", captureHandler)
	s.router.Post("/apiv1/transaction/refund", refundHandler)

	return http.ListenAndServe(addr, s.router)
}

// New return a new instance of the API service
func New() *Service {
	s := new(Service)
	return s
}
