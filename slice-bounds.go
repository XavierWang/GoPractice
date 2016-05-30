package main

import "fmt"

func main() {
	S := []int{2, 3, 5, 7, 11, 13}

	S = S[1:4]
	fmt.Println(S)

	S = S[:2]
	fmt.Println(S)

	S = S[1:]
	fmt.Println(S)
}
