package utils

const (
	MNT = "MNT"
	USD = "USD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case MNT, USD, EUR:
		return true
	}
	return false
}
