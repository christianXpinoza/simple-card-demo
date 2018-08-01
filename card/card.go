package card

import (
	"errors"
	"log"
	"sync"

	"github.com/bxcodec/faker"
)

// Card Abstraction
type Card struct {
	ID         uint64  `json:"id"`
	Number     string  `json:"number"`
	NameOnCard string  `json:"name"`
	Balance    float64 `json:"balance"`
}

// Storage struct for Cards
type Storage struct {
	sync.RWMutex
	Cards map[uint64]*Card
}

func (s *Storage) New(name string) (*Card, error) {
	var card Card

	newCard := struct {
		Number string `faker:"cc_number"`
	}{}

	s.Lock()
	defer s.Unlock()

	err := faker.FakeData(&newCard)
	if err != nil {
		log.Println("error generating card number")
		return &card, err
	}

	id := uint64(len(s.Cards)) + 1
	s.Cards[id] = &Card{
		ID:         id,
		Number:     newCard.Number,
		NameOnCard: name,
		Balance:    0,
	}

	return s.Cards[id], nil
}

// GetBalance ..
func (s *Storage) GetBalance(id uint64) (float64, error) {
	if c, ok := s.Cards[id]; ok {
		return c.Balance, nil
	}
	return 0, errors.New("client doesn't exist")
}

// Deposit ..
func (s *Storage) Deposit(id uint64, amount float64) (float64, error) {

	s.Lock()
	defer s.Unlock()

	if c, ok := s.Cards[id]; ok {
		c.Balance += amount
		return c.Balance, nil
	}

	return 0, errors.New("client doesn't exist")
}

// EarMark ...
/*
func (b *Storage) EarMark(id uint64, amount float64) (float64, error) {

	b.Lock()
	if bal, ok := b.Cards[id]; ok {
		bal.EarMarkedAmount += amount
		return bal.EarMarkedAmount, nil
	}
	b.Unlock()

	return 0, errors.New("client doesn't exist")
}
*/
