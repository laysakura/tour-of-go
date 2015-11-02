// This program does not work in localhost.
// Paste onto https://go-tour-jp.appspot.com/#36

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	retImg := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		retImg[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			//retImg[y][x] = uint8(x^y)
			//retImg[y][x] = uint8((x + y) / 2)
			retImg[y][x] = uint8(x * y)
		}
	}
	return retImg
}

func main()  {
	pic.Show(Pic)
}
