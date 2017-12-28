package main

import "fmt"

/**
 *
 * for 是 Go 的 “while”
 * 基于此可以省略分号：C 的 while 在 Go 中叫做 for 。
 */

//https://tour.go-zh.org/flowcontrol/3
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
