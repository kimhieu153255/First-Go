package main

import (
	"fmt"

	gobasics "github.com/kimhieu/first-go/cmd/go-basics"
)

func main() {
	gobasics.ArrayTest()

	var a [4]int
	for i := 0; i < 4; i++ {
		a[i] = i
	}

	fmt.Println(a)
}
