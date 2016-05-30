package main

import "golang.org/x/tour/pic"
import (
	"image"
	"image/color"
)

type Image struct {
	w, h   int
	mSlice [][]byte
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.w, m.h)
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{m.mSlice[x][y], m.mSlice[x][y], 255, 255}
}

func main() {
	haha := make([][]byte, 100)
	for i := range haha {
		haha[i] = make([]byte, 100)
	}

	for i := range haha {
		for j := range haha[i] {
			haha[i][j] = byte(i * j * i)
		}
	}
	m := Image{100, 100, haha}
	pic.ShowImage(m)
}
