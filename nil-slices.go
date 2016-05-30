package main

import "fmt"

func main() {
	var S []int
	fmt.Println(S, len(S), cap(S))
	if S == nil {
		fmt.Println("nil!")
	}
}
