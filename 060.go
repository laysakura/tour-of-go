package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
	"image/color"
)

type Image struct{}

func (m Image) ColorModel() color.Model {
	return color.RGBA64Model
}
func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}
func (m Image) At(x, y int) color.Color {
	return color.RGBA{uint8(255 * x * y / (100 * 100)), 0, 0, 0}
}

func main()  {
	m := Image{}
	pic.ShowImage(m)
}
