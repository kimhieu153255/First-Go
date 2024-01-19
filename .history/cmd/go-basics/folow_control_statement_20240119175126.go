package gobasics

import "fmt"

//
func CountSum(last int) int {
	for i := 0; i <= last; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
