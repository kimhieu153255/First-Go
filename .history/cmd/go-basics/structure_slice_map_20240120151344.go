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

func (s *Student) ToS() string {
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)
}

// Array
func ArrayTest() {
	var a [5][4]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			a[i][j] = i + j
		}
	}
	fmt.Println(a)
}

// Slice
func SliceTest() {
	a := make([]int, 5)
	fmt.Println(a)
}

// append
// signatures: func append(s []T, vs ...T) []T

// range
// for i, v := range a | for _, v := range a | for i := range a

// map
// signatures: map[K]V
// check: v, ok := m["Answer"]
