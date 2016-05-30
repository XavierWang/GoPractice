package main

import "fmt"

func main() {
	S := []int{2, 3, 5, 7, 11, 13}
	printSlice(S)

	S = S[:0]
	printSlice(S)

	S = S[:4]
	printSlice(S)

	S = S[3:]
	printSlice(S)

}

func printSlice(S []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(S), cap(S), S)
}
