package main

import "fmt"
/**
 * 一个结构体（ struct ）就是一个字段的集合。
 *（而 type 的含义跟其字面意思相符。）
 */

type Vertex struct {
	X int
	Y int
}

// https://tour.go-zh.org/moretypes/2
func main() {
	fmt.Println(Vertex{1, 2})
}
