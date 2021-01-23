package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

/*

指针与函数
现在我们要把 Abs 和 Scale 方法重写为函数。

同样，我们先试着移除掉第 16 的 *。你能看出为什么程序的行为改变了吗？要怎样做才能让该示例顺利通过编译？

（若你不确定，继续往下看。）
*/
func Abs1(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs1(v))
}
