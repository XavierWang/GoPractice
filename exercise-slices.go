package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	Plot := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		Plot[i] = make([]uint8, dx)
	}

	for i, row := range Plot {
		for j := range row {
			Plot[i][j] = uint8(i ^ j)
		}
	}
	return Plot
}

func main() {

	pic.Show(Pic)
}
