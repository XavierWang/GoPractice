package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (x MyFloat) Abs() float64 {
	if x < 0 {
		return float64(-x)
	}
	return float64(x)
}

func main() {
	x := MyFloat(-math.Sqrt(2))
	fmt.Println(x.Abs())

}
