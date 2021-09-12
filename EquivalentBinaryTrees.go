package main

import ("golang.org/x/tour/tree"
		"fmt"
	   )
		

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if(t == nil){
		return 
	}
	if(t.Left != nil){
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if(t.Right != nil){
		Walk(t.Right,ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	
	for i:=0; i < 10 ; i ++{
		x, y := <-c1, <- c2
		
		if x!=y{
			return false
		}
		
	}
	return true
}

func main() {
	tree1:= tree.New(9)
	tree2:= tree.New(9)
	fmt.Println(tree1)
	fmt.Println(tree2)
	
	fmt.Println(Same(tree1, tree2) )
}
