package main

import "fmt"

/*
 *
 * 变量在定义时没有明确的初始化时会赋值为 零值 。
 * 零值是：
 * 数值类型为 0 ，
 * 布尔类型为 false ，
 * 字符串为 "" （空字符串）。
 **/
//https://tour.go-zh.org/basics/12
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
