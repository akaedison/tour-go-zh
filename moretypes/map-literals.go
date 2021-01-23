package main

import "fmt"

type VertexMapLiterals struct {
	Lat, Long float64
}

/*
映射的文法
映射的文法与结构体相似，不过必须有键名。
*/
var ma = map[string]VertexMapLiterals{
	"Bell Labs": VertexMapLiterals{
		40.68433, -74.39967,
	},
	"Google": VertexMapLiterals{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(ma)
}
