package main

import "fmt"

func fibonacci() func() int {
	fibNow := 0
	fibNext := 1
	return func() int {
		output := fibNow
		fibNow = fibNext
		fibNext = output + fibNow

		return output
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
