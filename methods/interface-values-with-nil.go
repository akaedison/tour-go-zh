package main

import "fmt"

type I2 interface {
	M()
}

type T2 struct {
	S string
}

func (t *T2) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

/*

底层值为 nil 的接口值
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。

注意: 保存了 nil 具体值的接口其自身并不为 nil。
*/
func main() {
	var i I2

	var t *T2
	i = t
	describe1(i)
	i.M()

	i = &T2{"hello"}
	describe1(i)
	i.M()
}

func describe1(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
