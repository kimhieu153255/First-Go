package gogenerics

func Check[T comparable](a, b T) bool {
	return a == b
}


