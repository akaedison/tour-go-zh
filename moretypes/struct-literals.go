package main

import "fmt"

type VertexL struct {
	X, Y int
}

/*
结构体文法
结构体文法通过直接列出字段的值来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。
*/
var (
	v1 = VertexL{1, 2}  //创建一个Vertex类型的结构体
	v2 = VertexL{X: 1}  //Y: 0 被隐式地赋予
	v3 = VertexL{}      //	X :0 Y: 0
	p  = &VertexL{1, 2} //创建一个*Vertex类型的结构体(指针)
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
