package gobasics

import "fmt"

//pointer:
func PointerTest() {
	i := 1
	z := &i
	y := &z
	fmt.Println(**y)
}

//structure:
type Student struct {
	Name string
	Age  int
}

