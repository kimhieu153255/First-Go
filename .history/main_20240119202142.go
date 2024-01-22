package main

import (
	"fmt"

	gobasics "github.com/kimhieu/first-go/cmd/go-basics"
)

func main() {
	gobasics.ArrayTest()

	var b []int = []int{1, 2, 3, 4}

	fmt.Println(len(b))

}
