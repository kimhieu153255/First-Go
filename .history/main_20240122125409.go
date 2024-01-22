package main

import (
	"fmt"

	practiceleetcode "github.com/kimhieu/first-go/cmd/practice_leetcode"
)

func main() {
	median := practiceleetcode.FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(median)
}
