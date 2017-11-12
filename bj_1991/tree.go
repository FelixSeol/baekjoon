// tree
package main

import (
	"fmt"
)

type Node struct {
	value byte
	left  *Node
	right *Node
}

var text = `A B C
B D .
C E F
E . .
F . G
D . .
G . . `

func search(tar byte, node *Node) *Node {
	var result *Node

	if node != nil {
		if node.value == tar {
			result = node
		}
		if result == nil {
			result = search(tar, node.left)
		}
		if result == nil {
			result = search(tar, node.right)
		}
	} else {

	}
	return result
}

func preorder(node *Node) {
	if node.value == '.' {
		return
	}
	fmt.Printf("%c", node.value)
	preorder(node.left)
	preorder(node.right)
}

func postorder(node *Node) {
	if node.value == '.' {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Printf("%c", node.value)
}

func inorder(node *Node) {
	if node.value == '.' {
		return
	}
	inorder(node.left)
	fmt.Printf("%c", node.value)
	inorder(node.right)
}
func main() {
	var t int
	root := new(Node)
	fmt.Scanln(&t)

	f := func(v byte, l byte, r byte) {
		if root.value == 0 {
			root.value = v
			root.left = &Node{l, nil, nil}
			root.right = &Node{r, nil, nil}
		} else {
			node := search(v, root)
			if node != nil {
				node.left = &Node{l, nil, nil}
				node.right = &Node{r, nil, nil}
			}
		}
	}
	for i := 0; i < t; i++ {
		var v, l, r byte
		re, _ := fmt.Scanf("%c %c %c\n", &v, &l, &r)
		f(v, l, r)
		fmt.Println(v, l, r)
		_ = re
	}
	preorder(root)
	fmt.Println()
	inorder(root)
	fmt.Println()
	postorder(root)
}
