package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walkImpl1(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walkImpl1(t.Left, ch, quit)
	select {
	case ch <- t.Value:
	case <-quit:
		return
	}
	walkImpl1(t.Right, ch, quit)
}

func Walk1(t *tree.Tree, ch, quit chan int) {
	walkImpl1(t, ch, quit)
	close(ch)

}

func Same1(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk1(t1, w1, quit)
	go Walk1(t2, w2, quit)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
