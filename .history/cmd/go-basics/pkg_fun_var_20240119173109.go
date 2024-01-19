package gobasics

import (
	"fmt"
)

// 1. exported function: hàm này có thể được gọi từ bên ngoài package (Chữ đầu tiên của tên hàm phải viết hoa)

func countSum(a, b int) int {
	return a + b
}

func PkgFunVar() {
	fmt.Println("you are in pkgFunVar")

	var a int = 1
	var b int = 2

	fmt.Println(countSum(a, b))
}
