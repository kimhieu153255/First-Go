package gobasics // this is a package declaration

import (
	"fmt"
)

func PkgFunVar() (string, string) {
	return "pkgFunVar1", "pkgFunVar1"
} // this is a function declaration with a return type

func countSum(a, b int) int {
	return a + b
}

func pkgFunVar() {
	fmt.Println("you are in pkgFunVar")

}
