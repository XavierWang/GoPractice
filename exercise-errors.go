package main

import (
	"fmt"
	"math"
)

type ErrNegtiveSqrt float64

func (e ErrNegtiveSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negtive number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegtiveSqrt(x)
	}

	Z := 2.0
	for Delta := 1.0; Delta > 1e-10; {
		ZNext := Z - (Z*Z-x)/(2*Z)
		Delta = math.Abs(Z - ZNext)
		Z = ZNext
	}

	return Z, nil

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
