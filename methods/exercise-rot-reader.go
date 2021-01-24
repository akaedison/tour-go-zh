package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

/*
练习：rot13Reader
有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。

例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。

编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。

rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。
*/
func (rot rot13Reader) Read(p []byte) (n int, err error) {
	length, err := rot.r.Read(p)
	if err != nil {
		return length, err
	}
	for i := 0; i < length; i++ {
		v := p[i]
		switch {
		case 'a' <= v && v <= 'm':
			fallthrough
		case 'A' <= v && v <= 'M':
			p[i] = v + 13
		case 'n' <= v && v <= 'z':
			fallthrough
		case 'N' <= v && v <= 'Z':
			p[i] = v - 13
		}
	}
	return length, nil
}

/*
You cracked the code!
*/
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
