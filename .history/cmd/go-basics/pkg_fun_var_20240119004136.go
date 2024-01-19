package gobasics

import (
	"fmt"
)

func countSum(a, b int) int {
	return a + b
}

func pkgFunVar() {
	fmt.Println("you are in pkgFunVar")

	var a int = 1
	var b int = 2

	fmt.Println(countSum(a, b))
}
