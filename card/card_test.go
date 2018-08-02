package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() *Service {
	cardService := New()
	cardService.Storage.New("mark phantom")
	return cardService
}

func TestNewStorage(t *testing.T) {
	cardService := setup()
	// Create new Card
	card, err := cardService.Storage.New("peter parker")
	if assert.Nil(t, err) {
		assert.Equal(t, 0.0, card.Balance)
		assert.Equal(t, uint64(2), card.ID)
		assert.Equal(t, "peter parker", card.NameOnCard)
		assert.NotEqual(t, 0, card.Number)
		assert.NotNil(t, card.blockedAmounts.Amounts)
	}
}

func TestDepositGetBalance(t *testing.T) {
	cardService := setup()
	_, deposited, err := cardService.Storage.Deposit(1, 1000)
	assert.Nil(t, err)

	balance, blocked, err := cardService.Storage.GetBalance(1)
	if assert.Nil(t, err) {
		assert.Equal(t, balance, deposited)
		assert.NotEqual(t, 0.0, balance)
		assert.Equal(t, 0.0, blocked)
	}
}

func TestBlockAuthRequest(t *testing.T) {
	cardService := setup()
	cardService.Storage.Deposit(1, 1000)
	blockingID, err := cardService.Storage.BlockAuthRequest(1, 500)
	if assert.Nil(t, err) {
		assert.Equal(t, uint64(1), blockingID)
		balance, blocked, err := cardService.Storage.GetBalance(1)
		if assert.Nil(t, err) {
			assert.Equal(t, 500.0, balance)
			assert.Equal(t, 500.0, blocked)
		}
	}
	// Test cancel block auth
	err = cardService.Storage.CancelBlockAuth(1, blockingID)
	if assert.Nil(t, err) {
		balance, blocked, err := cardService.Storage.GetBalance(1)
		if assert.Nil(t, err) {
			assert.Equal(t, 1000.0, balance)
			assert.Equal(t, 0.0, blocked)
		}
	}
}

func TestCaptureBlockAuth(t *testing.T) {
	cardService := setup()
	cardID := uint64(len(cardService.Storage.Cards))
	cardService.Storage.Deposit(cardID, 1888)
	blockingID, err := cardService.Storage.BlockAuthRequest(cardID, 888)

	if assert.Nil(t, err) {
		_, captured, err := cardService.Storage.CaptureRequest(cardID, blockingID)
		if assert.Nil(t, err) {
			assert.Equal(t, 888.0, captured)
			balance, blocked, err := cardService.Storage.GetBalance(cardID)
			assert.Nil(t, err)
			assert.Equal(t, 1000.0, balance)
			assert.Equal(t, 0.0, blocked)
		}
	}

}

func TestRefunds(t *testing.T) {
	cardService := setup()

	cardID := uint64(len(cardService.Storage.Cards))
	cardService.Storage.Deposit(cardID, 1888)

	blockingID, err := cardService.Storage.BlockAuthRequest(cardID, 888)
	captureID, _, err := cardService.Storage.CaptureRequest(cardID, blockingID)

	_, err = cardService.Storage.Refund(cardID, captureID, 888)
	assert.Nil(t, err)

	balance, blocked, err := cardService.Storage.GetBalance(cardID)
	assert.Nil(t, err)
	assert.Equal(t, 1888.0, balance)
	assert.Equal(t, 0.0, blocked)
}
