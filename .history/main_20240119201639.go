package main

import (
	"fmt"

	gobasics "github.com/kimhieu/first-go/cmd/go-basics"
)

func main() {
	gobasics.ArrayTest()

	a := [4]int{1, 2, 3, 4}
	fmt.Println(a[1:3])

}
