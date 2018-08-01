package card

type Service struct {
	Storage      Storage
	Merchants    []Merchants
	Transactions chan Transaction
}

func New() *Service {
	var storage Storage
	storage.Cards = make(map[uint64]*Card)

	return &Service{
		Storage: storage,
	}
}
