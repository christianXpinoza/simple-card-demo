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

// Storage struct representation for Cards
type Storage struct {
	sync.RWMutex
	Cards        map[uint64]*Card
	Transactions *Transactions
}

// New Creates a new card
// Returns basic data from the new card and error
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
// Returns: Balance, Blocked Marked and error
func (s *Storage) GetBalance(id uint64) (float64, float64, error) {
	if c, ok := s.Cards[id]; ok {
		blockedTotal, err := c.blockedAmounts.GetTotal()
		if err != nil {
			return 0, 0, err
		}
		return c.Balance, blockedTotal, nil
	}
	return 0, 0, errors.New("card doesn't exist")
}

// Deposit adds an amount of £ to the balance of a card with id (id)
// Returns the current balance and error
func (s *Storage) Deposit(cardID uint64, amount float64) (uint64, float64, error) {

	// Protection for concurrent use of the map Cards
	s.Lock()
	defer s.Unlock()

	if c, ok := s.Cards[cardID]; ok {
		c.Balance += amount

		txnID := s.Transactions.Add(&Transaction{
			Kind:   "deposit",
			CardID: cardID,
			Amount: amount,
			Status: "done",
		})
		log.Println("deposit made, transaction id:", txnID)
		return txnID, c.Balance, nil
	}

	return 0, 0, errors.New("card doesn't exist")
}

// BlockAuthRequest blocks an amount of £ based in the card ID
// Returns the blocking ID to be used later to cancel or capture the £ & error
func (s *Storage) BlockAuthRequest(cardID uint64, amount float64) (uint64, error) {

	s.Lock()
	defer s.Unlock()

	if c, ok := s.Cards[cardID]; ok {

		// The merchant can Block the amount
		if c.Balance >= amount {
			blockID, err := c.blockedAmounts.Append(amount)
			if err != nil {
				return 0, err
			}
			// Blocked money is traspassed to blocked storing data structure
			c.Balance -= amount
			return blockID, nil
		}
		return 0, errors.New("not enough money")
	}

	return 0, errors.New("card doesn't exist")
}

// CaptureRequest try to capture an mount of blocked £ by card id & blocking id
// the amount is remove from blocked state
// Returns the amount captured an error
func (s *Storage) CaptureRequest(cardID, blockID uint64) (uint64, float64, error) {
	s.Lock()
	defer s.Unlock()

	if c, ok := s.Cards[cardID]; ok {
		if captured, ok := c.blockedAmounts.Amounts[blockID]; ok {
			delete(c.blockedAmounts.Amounts, blockID)

			txnID := s.Transactions.Add(&Transaction{
				Kind:   "capture",
				CardID: cardID,
				Amount: captured,
				Status: "done",
			})

			return txnID, captured, nil
		}
		return 0, 0, errors.New("blockID doesnt exist")
	}
	return 0, 0, errors.New("card doesn't exist")
}

// CancelBlockAuth cancels a Blocking Authotization by card id and blocking id
// Returns error
func (s *Storage) CancelBlockAuth(cardID, blockID uint64) error {
	s.Lock()
	defer s.Unlock()
	if c, ok := s.Cards[cardID]; ok {
		if captured, ok := c.blockedAmounts.Amounts[blockID]; ok {
			// returns the captured money to the available balance
			c.Balance += captured
			delete(c.blockedAmounts.Amounts, blockID)
			return nil
		}
		return errors.New("block id doesn't exist")
	}
	return errors.New("card doesn't exist")
}

// Refund will search in the transactions to match info and process refunding
func (s *Storage) Refund(cardID, captureID uint64, amount float64) (uint64, error) {
	for _, v := range s.Transactions.Transaction {
		if v.ID == captureID && v.CardID == cardID && v.Amount == amount {
			s.Deposit(cardID, amount)

			txnID := s.Transactions.Add(&Transaction{
				Kind:   "refund",
				CardID: cardID,
				Amount: amount,
				Status: "done",
			})
			return txnID, nil
		}

	}
	return 0, errors.New("error trying to process refund")
}
