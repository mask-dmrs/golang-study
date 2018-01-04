package main

import "fmt"

/**
 * defer 语句会延迟函数的执行直到上层函数返回。
 * 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
 */

//https://tour.go-zh.org/flowcontrol/12
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
