package gobasics // this is a package declaration

import (
	"fmt"
)

func PkgFunVar() (string, string) {
	return "pkgFunVar1", "pkgFunVar1"
}

func pkgFunVar() {
	fmt.Println("you are in pkgFunVar")

}
