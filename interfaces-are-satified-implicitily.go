package main

import "fmt"

type Ier interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i Ier
	i = T{"hello"}
	i.M()
}
