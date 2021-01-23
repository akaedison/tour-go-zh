package main

import "fmt"

/*
函数
函数可以没有参数或接受多个参数。

在本例中，add 接受两个 int 类型的参数。

注意类型在变量名 之后。

（参考 这篇关于 Go 语法声明(http://blog.go-zh.org/gos-declaration-syntax)的文章了解这种类型声明形式出现的原因。）
*/
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
