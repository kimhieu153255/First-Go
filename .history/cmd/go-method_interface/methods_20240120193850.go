package gomethodinterface

import "fmt"

// pointer:
type Student struct {
	Name string
	Age  int
}

// change value of a structure then use pointer
func (s *Student) ToS() string {
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)
}

// interface:
type Animal interface {
	Speak() string
	Run()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Gau gau"
}

func (d *Dog) Run() {
	fmt.Println("Dog is running")
}
