package main

import (
	"fmt"
	"runtime"
)

/*除非以 fallthrough 语句结束，否则分支会自动终止。*/

//https://tour.go-zh.org/flowcontrol/9
func main() {
	fmt.Print("Go runs on ")
	/*switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}*/
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("OS X.")
		fallthrough
	case "linux":
		fmt.Println("Linux.")
		fallthrough
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
