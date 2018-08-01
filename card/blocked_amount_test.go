package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupBlockAmountTest() *BlockedAmount {
	return &BlockedAmount{
		Amounts: make(map[uint64]float64),
	}
}

func TestBlockAmounts(t *testing.T) {
	blockedAmounts := setupBlockAmountTest()
	assert.IsType(t, &BlockedAmount{}, blockedAmounts)
	blockingID1, err := blockedAmounts.Append(55.21)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), blockingID1)

	blockingID2, err := blockedAmounts.Append(10.0)
	assert.Nil(t, err)
	assert.Equal(t, uint64(2), blockingID2)

	amount1, err := blockedAmounts.Get(blockingID1)
	assert.Nil(t, err)
	assert.Equal(t, 55.21, amount1)

	amount2, err := blockedAmounts.Get(blockingID2)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, amount2)

	total, err := blockedAmounts.GetTotal()
	assert.Nil(t, err)
	assert.Equal(t, amount1+amount2, total)

	nonExistingAmount, err := blockedAmounts.Get(3)
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, nonExistingAmount)
}
