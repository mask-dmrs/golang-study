package main

import "fmt"

func add(x int, y int) int {
	return x + y;
}

func printTest(x string) {
	fmt.Print(x);
}

func main() {
	fmt.Println(add(42, 13));
	printTest("XXXX");
}
