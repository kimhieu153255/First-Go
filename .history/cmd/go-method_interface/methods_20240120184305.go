package gomethodinterface

import "fmt"

// pointer:

// structure:
type Student struct {
	Name string
	Age  int
}

// change value of a structure then use pointer
func (s *Student) ToS() string {
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)
}


