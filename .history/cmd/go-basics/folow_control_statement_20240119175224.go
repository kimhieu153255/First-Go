package gobasics

// for bắt buộc phải có {}
func CountSum(last int) int {
	result := 0
	for i := 0; i <= last; i++ {
		result += i
	}
	return result
}
