package gobasics // this is a package declaration

import (
	"fmt"
)

func countSum(a, b int) int {
	return a + b
}

func PkgFunVar() {
	fmt.Println("you are in pkgFunVar")

	var a int = 1
	var b int = 2

	fmt.Println(countSum(a, b))
}

func init() {
	pkgFunVar()
}
