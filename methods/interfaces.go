package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

/*
接口
接口类型 是由一组方法签名定义的集合。

接口类型的变量可以保存任何实现了这些方法的值。

注意: 示例代码的 22 行存在一个错误。由于 Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser。
*/
func main() {
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := Vertex7{3, 4}

	a = f //a MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v //a *Vertex7 实现了 Abser
	fmt.Println(a.Abs())
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// a = v

}

type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex7 struct {
	X, Y float64
}

func (v *Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
