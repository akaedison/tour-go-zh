package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1 //初始值为0,1
	return func() int {
		a, b = b, a+b //将a变为b,b变为a+b
		return b - a  //此时应该返回为改变时的a, 也就是a+b-b = b-a
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
