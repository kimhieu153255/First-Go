package gobasics

import "fmt"

func PointerTest() {
	i := 1
	z := &i
	y := &z
	fmt.Println(**y)
}
