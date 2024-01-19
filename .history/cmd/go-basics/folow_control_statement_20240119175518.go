package gobasics

import "fmt"

// for bắt buộc phải có {}
func CountSum(last int) {
	result := 0
	for i := 0; i <= last; i++ {
		result += i
	}
	fmt.Println(result)
}

// for continued
func CountSumForContinued(last int) {
	result, ind := 0
	for ind <= last {

	}
}
