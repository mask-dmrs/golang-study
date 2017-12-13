package ch02

import "fmt"

func add(x int, y int) int {
	return x + y;
}

func printTest(x string) {
	fmt.Print(x);
}
//https://tour.go-zh.org/basics/4
func main() {
	fmt.Println(add(42, 13));
	printTest("XXXX");
}
