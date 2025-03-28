package headers

import (
	"github.com/steve-care-software/grammars/domain/roles/headers/prices"
	"github.com/steve-care-software/grammars/domain/roles/headers/taxes"
)

// Header represents the header
type Header interface {
	Version() uint
	HasPrices() bool
	Prices() prices.Prices
	HasTaxes() bool
	Taxes() taxes.Taxes
}
