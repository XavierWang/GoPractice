package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	Z := 2.0
	for Delta := 1.0; Delta > 1e-10; {
		ZNext := Z - (Z*Z-x)/(2*Z)
		Delta = math.Abs(Z - ZNext)
		Z = ZNext
	}

	return Z

}

func main() {
	fmt.Println(Sqrt(2))
}
