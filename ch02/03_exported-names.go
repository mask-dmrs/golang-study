package ch02

/**
 * 在 Go 中，首字母大写的名称是被导出的。
 *
 * 在导入包之后，你只能访问包所导出的名字，任何未导出的名字是不能被包外的代码访问的。
 *
 * Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。
 */

import (
	"fmt"
	"math"
)
//https://tour.go-zh.org/basics/3
func main() {
	fmt.Println(math.Pi);
}
