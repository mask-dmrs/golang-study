package main

import "fmt"

/*
 * 切片拥有 长度 和 容量 。
 *
 * 切片的长度就是它所包含的元素个数。
 *
 * 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
 *
 * 切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
 *
 * 你可以通过重新切片来扩展一个切片，给它提供足够的容量。 试着修改示例程序中的切片操作，向外扩展它的容量，看看会发生什么。
 */

//https://tour.go-zh.org/moretypes/11

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
