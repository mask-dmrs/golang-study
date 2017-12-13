package ch02

import (
	"fmt"
	"math/rand"
	"time"
)
//https://tour.go-zh.org/basics/1
//这个程序的运行环境是确定性的，因此 rand.Intn 每次都会返回相同的数字。 （为了得到不同的随机数，需要提供一个随机数种子，参阅 rand.Seed。）
func main() {
	rand.Seed(time.Now().Unix());
	fmt.Println("My favorite number is", rand.Intn(10));
}
