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

type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	return "Meo meo"
}

func (c *Cat) Run() {
	fmt.Println("Cat is running")
}

func (a Student) String() string {
	return fmt.Sprintf("Name: %q; Age: %d", a.Name, a.Age)
}

// stringer:
func TestStringer() {
	student := Student{
		Name: "Hieu",
		Age:  22,
	}
	fmt.Println(student)
}

// khởi tạo pointer
func PointerTest() {
	// z là con trở * int
	var z *int = new(int)
	*z = 1
}
