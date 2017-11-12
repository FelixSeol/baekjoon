package main

import "fmt"

var root Node

type Node struct{
	value byte
	left *Node
	right *Node
}

func traverse(tar byte) *Node{
	node := &root
	fmt.Print(node.value)
	return node
}


func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i<t ; i++{
		var v, l, r byte
		fmt.Scan(&v, &l, &r)
		if root.value == 0 {
			root.value = v
			root.left = &Node{l, nil, nil}
			root.right = &Node{r, nil, nil}
		}
	}

	fmt.Printf("root v : %v   root l : %v   root r : %v \n",root.value, root.left.value, root.right.value)
}
