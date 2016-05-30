package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	V1 = Vertex{1, 2}
	V2 = Vertex{X: 1}
	V3 = Vertex{}
	P  = &Vertex{1, 2}
)

func main() {
	fmt.Println(V1, P, V2, V3)
}
