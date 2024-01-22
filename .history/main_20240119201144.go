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
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			a[i][j] = i + j
		}
	}
	fmt.Println(a)
}
