package taxes

const (
	// TAX_REFERRAL represents a referral tax
	TAX_REFERRAL (uint8) = iota

	// TAX_DAO represents the dao tax
	TAX_DAO
)

// Taxes represents taxes
type Taxes interface {
	List() []Tax
}

// Tax represents a tax
type Tax interface {
	Kind() uint8
	Percent() float32
}
