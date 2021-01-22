package main

import "fmt"

type VertexF struct {
	X int
	Y int
}
/*

结构体字段
结构体字段使用点号来访问。
*/
func main() {
	v := VertexF{1,2}
	v.X = 4
	fmt.Println(v.X)
}
