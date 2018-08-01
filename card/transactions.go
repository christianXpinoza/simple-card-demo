package card

import (
	"time"
)

// Transaction representation
type Transaction struct {
	ID         uint64    `json:"id"`          // Unique Transaction ID
	Kind       string    `json:"kind"`        // Transaction Kind []
	Datetime   time.Time `json:"datetime"`    // UTC date time
	Amount     float64   `json:"amount"`      // amount of £ involved
	Status     string    `json:"status"`      // Transaction status []
	ClientID   uint64    `json:"client_id"`   // Client ID
	MerchantID uint64    `json:"merchant_id"` // Merchant ID
}

// Representation of transactions
type Transactions chan Transaction
