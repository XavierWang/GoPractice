package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (y MyFloat) Abs() float64 {
	if y < 0 {
		return float64(y)
	}
	return float64(y)
}

type Vertex struct {
	X, Y float64
}

func (y *Vertex) Abs() float64 {
	return math.Sqrt(y.X*y.X + y.Y*y.Y)

}
func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}
