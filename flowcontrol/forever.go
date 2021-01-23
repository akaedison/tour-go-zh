package main

import "fmt"

/*
无限循环
如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
*/
func main() {
	for {
		fmt.Println("Hello")
	}
}
