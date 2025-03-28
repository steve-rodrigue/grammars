package prices

// Prices represents prices
type Prices interface {
	List() []Price
}

// Price represents a price
type Price interface {
	Permission() string
	Amount() uint64
}
