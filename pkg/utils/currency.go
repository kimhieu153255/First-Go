package utils

var (
	USD = "USD"
	EUR = "EUR"
	VND = "VND"
)

func IsCurrency(currency string) bool {
	switch currency {
	case USD, EUR, VND:
		return true
	}
	return false
}
