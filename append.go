package main

import "fmt"

func main() {
	var S []int

	printSlice(S)

	S = append(S, 0)
	printSlice(S)

	S = append(S, 1)
	printSlice(S)

	S = append(S, 2, 3, 4, 5, 6)
	printSlice(S)
}

func printSlice(S []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(S), cap(S), S)

}
