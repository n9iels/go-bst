package main

import "fmt"

// Node for a BinarySearchTree
type Node struct {
	Key                 int
	Value               string
	Left, Right, Parent *Node
}

// Binary Search Tree
type BinarySearchTree struct {
	Root *Node
}

// Internal insert method
func (bts BinarySearchTree) Insert(root *Node, key int, value string) {
	var newNode = &Node{key, value, nil, nil, root}

	if root.Left == nil && key < root.Key {
		root.Left = newNode
	} else if root.Right == nil && key > root.Key {
		root.Right = newNode
	} else if key < root.Key {
		bts.Insert(root.Left, key, value)
	} else {
		bts.Insert(root.Right, key, value)
	}
}

// Internal traversal method
func (bts *BinarySearchTree) Traverse(node *Node, key int) *Node {
	if key > node.Key {
		return bts.Traverse(node.Right, key)
	} else if key < node.Key {
		return bts.Traverse(node.Left, key)
	}

	return node
}

// External add method
func (bts *BinarySearchTree) Add(key int, value string) {
	if bts.Root == nil {
		bts.Root = &Node{key, value, nil, nil, nil}
	} else {
		bts.Insert(bts.Root, key, value)
	}
}

// External search method
func (bts *BinarySearchTree) Search(key int) *Node {
	return bts.Traverse(bts.Root, key)
}

func main() {
	var bts = BinarySearchTree{}
	bts.Add(6, "six")
	bts.Add(5, "five")
	bts.Add(7, "seven")
	bts.Add(3, "three")
	bts.Add(1, "one")

	fmt.Println(bts.Search(1).Value)
}
