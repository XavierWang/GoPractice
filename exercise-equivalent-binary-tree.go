package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var walk func(t *tree.Tree, ch chan int)
	walk = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
	walk(t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch_t1 := make(chan int)
	ch_t2 := make(chan int)
	go Walk(t1, ch_t1)
	go Walk(t2, ch_t2)

	for node := range ch_t1 {
		if node != <-ch_t2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(10)
	t2 := tree.New(10)
	fmt.Println(Same(t1, t2))
}
