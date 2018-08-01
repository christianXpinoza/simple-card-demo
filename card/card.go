package card

import (
	"errors"
	"log"
	"sync"

	"github.com/bxcodec/faker"
)

// Card Abstraction
type Card struct {
	ID             uint64        `json:"id"`
	Number         string        `json:"number"`
	NameOnCard     string        `json:"name"`
	Balance        float64       `json:"balance"`
	blockedAmounts BlockedAmount // non-exported field
}

// Storage struct for Cards
type Storage struct {
	sync.RWMutex
	Cards map[uint64]*Card
}

// New Create a new card
func (s *Storage) New(name string) (*Card, error) {
	var card Card

	newCard := struct {
		Number string `faker:"cc_number"`
	}{}

	// Protection for concurrent use of the map Cards
	s.Lock()
	defer s.Unlock()

	// Generate new fake card number
	err := faker.FakeData(&newCard)
	if err != nil {
		log.Println("error generating card number")
		return &card, err
	}
	// Add new card to the storage
	id := uint64(len(s.Cards)) + 1
	s.Cards[id] = &Card{
		ID:             id,
		Number:         newCard.Number,
		NameOnCard:     name,
		Balance:        0,
		blockedAmounts: BlockedAmount{Amounts: make(map[uint64]float64)},
	}

	return s.Cards[id], nil
}

// GetBalance return the balance for a client with card id (id)
func (s *Storage) GetBalance(id uint64) (float64, error) {
	if c, ok := s.Cards[id]; ok {
		return c.Balance, nil
	}
	return 0, errors.New("client doesn't exist")
}

// Deposit add amount of £ to the balance of a card with id (id)
func (s *Storage) Deposit(id uint64, amount float64) (float64, error) {

	// Protection for concurrent use of the map Cards
	s.Lock()
	defer s.Unlock()

	if c, ok := s.Cards[id]; ok {
		c.Balance += amount
		return c.Balance, nil
	}

	return 0, errors.New("client doesn't exist")
}

// BlockAuthRequest ...
func (s *Storage) BlockAuthRequest(cardID uint64, amount float64) (uint64, error) {

	s.Lock()
	defer s.Unlock()

	if c, ok := s.Cards[cardID]; ok {
		totalBlocked, err := c.blockedAmounts.GetTotal()
		if err != nil {
			return 0, err
		}
		// The merchant can Block the amount
		if c.Balance-totalBlocked >= amount {
			blockID, err := c.blockedAmounts.Append(amount)
			if err != nil {
				return 0, err
			}
			return blockID, nil
		}
		return 0, errors.New("not enough money")
	}

	return 0, errors.New("card doesn't exist")
}
