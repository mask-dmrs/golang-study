package main

import (
	"fmt"
	"time"
)

/**
 * 没有条件的 switch
 * 没有条件的 switch 同 switch true 一样。
 *
 * 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
 */

//https://tour.go-zh.org/flowcontrol/11
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
