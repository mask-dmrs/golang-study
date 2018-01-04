package main

import (
	"fmt"
	"time"
)

/**
 * switch 的执行顺序
 * switch 的条件从上到下的执行，当匹配成功的时候停止。
 * （例如，
 * switch i {
 * case 0:
 * case f():
 *  }
 * 当 i==0 时不会调用 f 。）
 * 注意：Go playground 中的时间总是从 2009-11-10 23:00:00 UTC 开始， 如何校验这个值作为一个练习留给读者完成。
 */

//https://tour.go-zh.org/flowcontrol/10
func main() {
	fmt.Println("When's Saturday?")
	//当在web执行以下语句打印这个时间：  2009-11-10 23:00:00 +0000 UTC m=+0.000000001
	fmt.Println(time.Now())
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
