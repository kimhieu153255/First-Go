package main // package name must be main if you want to run this file as a program

import (
	"fmt"

	gobasics "github.com/kimhieu/first-go/cmd/go-basics"
)

func main() {
	s := gobasics.Student{Name: "Hieu", Age: 22}
	fmt.Println(s.ToS())
}

// Array
func ArrayTest() {
	var a [5][4]int
	a[2] = 1
	fmt.Println(a)
}
