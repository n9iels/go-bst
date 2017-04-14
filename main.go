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
func (bst BinarySearchTree) Insert(root *Node, key int, value string) {
	var newNode = &Node{key, value, nil, nil, root}

	if root.Left == nil && key < root.Key {
		root.Left = newNode
	} else if root.Right == nil && key > root.Key {
		root.Right = newNode
	} else if key < root.Key {
		bst.Insert(root.Left, key, value)
	} else {
		bst.Insert(root.Right, key, value)
	}
}

// Internal traversal method
func (bst *BinarySearchTree) Traverse(node *Node, key int) *Node {
	if key > node.Key {
		return bst.Traverse(node.Right, key)
	} else if key < node.Key {
		return bst.Traverse(node.Left, key)
	}

	return node
}

// External add method
func (bst *BinarySearchTree) Add(key int, value string) {
	if bst.Root == nil {
		bst.Root = &Node{key, value, nil, nil, nil}
	} else {
		bst.Insert(bst.Root, key, value)
	}
}

// External search method
func (bst *BinarySearchTree) Search(key int) *Node {
	return bst.Traverse(bst.Root, key)
}

func main() {
	var bst = BinarySearchTree{}
	bst.Add(6, "six")
	bst.Add(5, "five")
	bst.Add(7, "seven")
	bst.Add(3, "three")
	bst.Add(1, "one")

	fmt.Println(bst.Search(1).Value)
}
