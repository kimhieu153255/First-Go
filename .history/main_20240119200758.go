package main // package name must be main if you want to run this file as a program

import (
	"fmt"

	gobasics "github.com/kimhieu/first-go/cmd/go-basics"
)

func main() {
	gobasics.StructureTest()
	s := gobasics.Student{"Hieu", 22}
	fmt.Println(s.ToS())
}
