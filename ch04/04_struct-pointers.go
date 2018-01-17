package main

/**
 * 结构体字段可以通过结构体指针来访问。
 *
 * 通过指针间接的访问是透明的。
 */

import "fmt"

type Vertex struct {
	X int
	Y int
}

// https://tour.go-zh.org/moretypes/4
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
