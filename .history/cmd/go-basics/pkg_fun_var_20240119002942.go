package gobasics // this is a package declaration

import (
	"fmt"
)

func countSum(a, b int) int {
	return a + b
}

func pkgFunVar() {
	fmt.Println("you are in pkgFunVar")
	// function call
	fmt.Println(countSum(1, 2))
}
