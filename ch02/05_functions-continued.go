package ch02

import "fmt"

func add(x, y int) int {
	return x + y;
}
//https://tour.go-zh.org/basics/5
func main() {
	fmt.Println(add(42, 13));
}
