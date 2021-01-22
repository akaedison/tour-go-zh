package main

import "fmt"
/*

变量
var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。
*/
var c, python, java bool
/*
0 false false false
*/
func main() {
	var i int
	fmt.Println(i, c, python, java)
}
