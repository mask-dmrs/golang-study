package main

import "fmt"

/**
 * 循环初始化语句和后置语句都是可选的。
 */

 //https://tour.go-zh.org/flowcontrol/2
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}