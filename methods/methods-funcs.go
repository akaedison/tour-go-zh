package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

/*

方法即函数
记住：方法只是个带接收者参数的函数。

现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
*/
func Abs(v Vertex1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex1{3, 4}
	fmt.Println(Abs(v))
}
