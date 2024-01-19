package gobasics // this is a package declaration

import (
	"fmt"
)

func PkgFunVar() (string, string) {
	return "pkgFunVar1", "pkgFunVar1"
} // this is a function declaration with a return type

func pkgFunVar1() {
	fmt.Println("you are in pkgFunVar1")

}

func pkgFunVar() {
	fmt.Println("you are in pkgFunVar")

}
