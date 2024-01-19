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
	var a int = 1
	var b int = 2

	fmt.Println(countSum(a, b))

}
