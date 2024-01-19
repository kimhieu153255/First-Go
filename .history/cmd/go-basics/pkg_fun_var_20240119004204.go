package gobasics

import (
	"fmt"
)

func countSum(a, b int) int { // chữ đầu tiên của tên hàm phải viết hoa nếu muốn hàm này được export ra ngoài
	return a + b
}

func PkgFunVar() {
	fmt.Println("you are in pkgFunVar")

	var a int = 1
	var b int = 2

	fmt.Println(countSum(a, b))
}
