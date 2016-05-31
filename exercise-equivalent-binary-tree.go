package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		close(ch)
		return
	}
	ch_left := make(chan int)
  ch_rigth := make(chan int)
  go Walk(t.Left, ch_left)
  go Walk(t.Right, ch_right)
  for left_finished,right_finished := false,false {
    select{
    case left_num,ok := <- ch_left:
      if ok {
        ch <- left_num
      }else {
        left_finished = true
      }
    case right_num,ok := <= ch_rigth:
      if ok {
        ch <- right_num
      }else{
        right_finished = true
      }
    default:
      if left_finished && right_finished {
        break
      }
    }
  }
  ch <- t.Value
  close(ch)
  return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
  tree1 := make(map[int]int)
  tree2 := make(map[int]int)
  ch_t1 = make(chan int)
  ch_t2 = make(chan int)
  go Walk(t1, ch_t1)
  go Walk(t2, ch_t2)
  for t1_finished,t2_finished := false,false {
    select{
    case t1_num,ok := <- ch_t1:
      if ok {
        tree1[t1_num]++
      }else {
        t1_finished = true
      }
    case t2_num,ok := <= ch_t2:
      if ok {
        tree2[t2_num]++
      }else{
        t2_finished = true
      }
    default:
      if t1_finished && t2_finished {
        break
      }
    }
  }

  if(tree1 == tree2)
    return true
  return false
}

func main() {
  t1 := tree.New(10)
  t2 := tree.New(10)
  fmt.Println(Same(t1, t2))
}
