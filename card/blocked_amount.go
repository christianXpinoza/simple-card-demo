package card

import (
	"errors"
	"sync"
)

// BlockedAmount is a structure containing the blocked amounts of £
type BlockedAmount struct {
	sync.RWMutex
	Amounts map[uint64]float64
}

// Append adds a new blocked amount of £
// Returns the index/id of the new blocked amount of £ & error
func (b *BlockedAmount) Append(newBlockedAmount float64) (uint64, error) {
	b.Lock()
	defer b.Unlock()

	n := uint64(len(b.Amounts))
	b.Amounts[n+1] = newBlockedAmount

	return n + 1, nil
}

// Delete removes a existing blocked amount of £
// Returns error
func (b *BlockedAmount) Delete(blockedAmountID uint64) error {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.Amounts[blockedAmountID]; ok {
		delete(b.Amounts, blockedAmountID)
		return nil
	}
	return errors.New("item doesn't exist")
}

// Get retrieves a blocked amount of £
// Returns the blocked amount of £ for a existing blocking id
func (b *BlockedAmount) Get(blockedAmountID uint64) (float64, error) {
	b.Lock()
	defer b.Unlock()

	if blockedAmount, ok := b.Amounts[blockedAmountID]; ok {
		return blockedAmount, nil
	}
	return 0, errors.New("item doesn't exist")
}

// GetTotal sums the total value stored as a blocked amount of £
// Returns total amount of blocked £ & error
func (b *BlockedAmount) GetTotal() (float64, error) {
	b.Lock()
	defer b.Unlock()

	var total float64

	for _, blocked := range b.Amounts {
		total += blocked
	}
	return total, nil
}
