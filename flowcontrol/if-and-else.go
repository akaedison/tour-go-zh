package main

import (
	"fmt"
	"math"
)

/*

if 和 else
在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

（在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。）
*/
func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		Pow(3, 2, 10),
		Pow(3, 3, 20),
	)
}
