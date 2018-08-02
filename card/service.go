package card

// Service structure
type Service struct {
	Storage   Storage     // Main Data Repository
	Merchants []Merchants // I was thinking in include a sort of auth but seems out of scope
}

// New create a new instance of card service
func New() *Service {
	var storage Storage
	storage.Cards = make(map[uint64]*Card)
	storage.Transactions = NewTransactionInstance()

	return &Service{
		Storage: storage,
	}
}
