package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

/*
练习：映射
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
*/
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		_, ok := m[word]
		if ok == false {
			m[word] = 1
		} else {
			m[word]++
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
