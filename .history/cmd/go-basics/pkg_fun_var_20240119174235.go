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

// zero value: giá trị mặc định của 1 biến khi khai báo mà không gán giá trị
// 1. bool: false
// 2. numeric: 0
// 3. string: ""
// 4. pointer: nil
// 5. function: nil
// 6. interface: nil
// 7. slice: nil
// 8. channel: nil
// 9. map: nil
