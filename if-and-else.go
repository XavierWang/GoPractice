package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if V := math.Pow(x, n); V < lim {
		return V
	} else {
		fmt.Printf("%g >= %g\n", V, lim)

	}

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
