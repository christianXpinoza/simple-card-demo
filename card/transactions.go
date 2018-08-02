package card

import (
	"errors"
	"sync"
	"time"
)

// Transaction representation
type Transaction struct {
	ID         uint64    `json:"id"`          // Unique Transaction ID
	Kind       string    `json:"kind"`        // Basic Transaction Kinds [deposit, capture, Refund]
	Datetime   time.Time `json:"datetime"`    // UTC date time
	Amount     float64   `json:"amount"`      // amount of £ involved
	Status     string    `json:"status"`      // Transaction status [pending, done]
	CardID     uint64    `json:"card_id"`     // Client ID
	MerchantID uint64    `json:"merchant_id"` // Merchant ID
	payload    []byte    // Data payload - no exported just for auditing purposes
}

// Transactions Log live representation
type Transactions struct {
	sync.RWMutex
	Transaction map[uint64]*Transaction
}

// NewTransactionInstance represent storage for transaction log
func NewTransactionInstance() *Transactions {
	return &Transactions{
		Transaction: make(map[uint64]*Transaction),
	}
}

// Add adds a new record
func (t *Transactions) Add(txn *Transaction) uint64 {
	t.Lock()
	defer t.Unlock()

	id := uint64(len(t.Transaction)) + 1
	txn.ID = id
	txn.Datetime = time.Now()
	txn.Status = "done"

	t.Transaction[id] = txn
	return id
}

// GetByCardID retrieves a slice of transaction without any limit at the moment by user ID
func (t *Transactions) GetByCardID(id uint64) ([]Transaction, error) {
	t.RLock()
	defer t.RUnlock()

	txns := []Transaction{}
	for _, v := range t.Transaction {
		if v.CardID == id {
			txns = append(txns, *v)
		}
	}
	if len(txns) > 0 {
		return txns, nil
	}
	return nil, errors.New("no existing transactions")
}

// GetByMerchantID retrieves a slice of transactions without any limit at the moment by merchant ID
func (t *Transactions) GetByMerchantID(id uint64) ([]Transaction, error) {
	t.RLock()
	defer t.RUnlock()

	txns := []Transaction{}
	for _, v := range t.Transaction {
		if v.MerchantID == id {
			txns = append(txns, *v)
		}
	}
	if len(txns) > 0 {
		return txns, nil
	}
	return nil, errors.New("no existing transactions")
}
