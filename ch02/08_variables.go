package ch02

import "fmt"

var c, python, java bool

/*
 * var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。
 * 就像在这个例子中看到的一样， var 语句可以定义在包或函数级别。
 **/
//https://tour.go-zh.org/basics/8
func main() {
	var i int
	fmt.Println(i,c,python,java);
}
