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

// declare a structure
func StructureTest() {
	student := Student{
		Name: "Hieu",
		Age:  22,
	}
	fmt.Println(student)
}

// 
