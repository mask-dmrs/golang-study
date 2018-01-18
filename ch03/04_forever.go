package main

import "fmt"

/**
 * 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
 */

//https://tour.go-zh.org/flowcontrol/4
func main() {
	for {
		fmt.Println("loop print test")
	}
}
