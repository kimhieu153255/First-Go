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
	result, ind := 0, 1
	for ind <= last {
		result += ind
		ind++
	}
	fmt.Println(result)
}

// while in go is for
func CountSumWhile(last int) {
	result, ind := 0, 1
	for ind <= last {
		result += ind
		ind++
	}
	fmt.Println(result)
}

// if bắt buộc phải có {}, giống bth không có ()
func CheckEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

// if with a short statement (count, condition)
func CheckEvenOddShortStatement(num int) {
	if v := num % 2; v == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

// switch case (có thể dùng
func CheckEvenOddSwitch(num int) {
	switch num % 2 {
	case 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}
}
