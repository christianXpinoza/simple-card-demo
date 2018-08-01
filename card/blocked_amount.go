package card

import (
	"errors"
	"sync"
)

// Blocked ...
//type Blocked struct {
//	Amount float64
//}

// BlockedAmount ...
type BlockedAmount struct {
	sync.RWMutex
	Amounts map[uint64]float64
}

// Append ...
func (b *BlockedAmount) Append(newBlockedAmount float64) (uint64, error) {
	b.Lock()
	defer b.Unlock()

	n := uint64(len(b.Amounts))
	b.Amounts[n+1] = newBlockedAmount

	return n + 1, nil
}

// Delete ...
func (b *BlockedAmount) Delete(blockedAmountID uint64) error {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.Amounts[blockedAmountID]; ok {
		delete(b.Amounts, blockedAmountID)
		return nil
	}
	return errors.New("item doesn't exist")
}

// Get ...
func (b *BlockedAmount) Get(blockedAmountID uint64) (float64, error) {
	b.Lock()
	defer b.Unlock()

	if blockedAmount, ok := b.Amounts[blockedAmountID]; ok {
		return blockedAmount, nil
	}
	return 0, errors.New("item doesn't exist")
}

// GetTotal ...
func (b *BlockedAmount) GetTotal() (float64, error) {
	b.Lock()
	defer b.Unlock()

	var total float64

	for _, blocked := range b.Amounts {
		total += blocked
	}
	return total, nil

}
