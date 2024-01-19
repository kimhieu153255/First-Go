package gobasics

import (
	"fmt"
)

// 1. exported: có thể được gọi từ bên ngoài package (Chữ đầu tiên của tên phải viết hoa)
// 2. unexported: chỉ có thể được gọi từ bên trong package
// 3. main package: là package chứa hàm main, chỉ có thể có 1 main package trong 1 project

func countSum(a, b int) int { // exported function
	return a + b
}

func PkgFunVar() {
	fmt.Println("you are in pkgFunVar")

	// var a, b int = 1, 2 | var a, b = 1, 2 | a, b := 1, 2
	var a int = 1
	var b int = 2

	fmt.Println(countSum(a, b))
}
