package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func _walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		_walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		_walk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	_walk(t,ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)
	//compare elements in the tree one by one 
	for i := range ch1 {
		fmt.Println("i is ", i)
		j := <-ch2
		fmt.Println("j is ", j)
		if i != j {
			return false
		}	
	}
	
	return true
}

func main() {
	t := tree.New(2)
	c := make(chan int, 10)
	go Walk(t, c)
	for i := range c {
		fmt.Println(i)
	}
	
	fmt.Println(Same(tree.New(2),tree.New(2)))	
}
