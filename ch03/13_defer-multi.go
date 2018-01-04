package main

import "fmt"

/**
 * 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
 *
 * 阅读博文（https://blog.go-zh.org/defer-panic-and-recover）了解更多关于 defer 语句的信息。
 */

//https://tour.go-zh.org/flowcontrol/13
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
