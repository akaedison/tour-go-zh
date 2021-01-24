package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

/*
练习：图像
还记得之前编写的图片生成器 吗？我们再来编写另外一个，不过这次它将会返回一个 image.Image 的实现而非一个数据切片。

定义你自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。

Bounds 应当返回一个 image.Rectangle ，例如 image.Rect(0, 0, w, h)。

ColorModel 应当返回 color.RGBAModel。

At 应当返回一个颜色。上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。


*/
type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	r := (uint8)((float64)(x) / (float64)(i.w) * 255.0)
	g := (uint8)((float64)(y) / (float64)(i.h) * 255.0)
	b := (uint8)((float64)(x*y) / (float64)(i.w*i.h) * 255.0)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
