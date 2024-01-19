package gobasics // this is a package declaration

import (
	"fmt"
)

func PkgFunVar() (string, string) {
	return "pkgFunVar", "pkgFunVar"
}

func pkgFunVar() {
	fmt.Println("you are in pkgFunVar")

}
