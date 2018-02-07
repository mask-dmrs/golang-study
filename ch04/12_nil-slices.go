package main

import "fmt"

/**
 * nil slice
 * slice 的零值是 nil 。
 *
 * 一个 nil 的 slice 的长度和容量是 0。
 */

//https://tour.go-zh.org/moretypes/12
func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
