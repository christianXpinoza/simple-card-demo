package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransactionInstance(t *testing.T) {
	trnInstance := NewTransactionInstance()
	assert.NotNil(t, trnInstance)
	assert.NotNil(t, trnInstance.Transaction)

}

func TestTransactions(t *testing.T) {
	trnInstance := NewTransactionInstance()
	txn := &Transaction{
		Kind:   "deposit",
		CardID: uint64(1),
		Amount: 1000.0,
		Status: "pending",
	}
	trnInstance.Add(txn)

	txns, err := trnInstance.GetByCardID(1)
	assert.Nil(t, err)
	assert.Len(t, txns, 1)

	txns, err = trnInstance.GetByMerchantID(100)
	assert.NotNil(t, err)
	assert.Len(t, txns, 0)
}
