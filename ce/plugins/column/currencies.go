package column

import "github.com/hanzo-fi/go-libs/v5/pkg/types/currency"

var (
	supportedCurrenciesWithDecimal = map[string]int{
		"USD": currency.ISO4217Currencies["USD"], // US Dollar
	}
)
