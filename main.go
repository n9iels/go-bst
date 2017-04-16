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

func (root *Node) isLeaf() bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}

// Internal insert method
// internal method are starting lowercase
func (root *Node) insert(key int, value string, parent *Node) *Node {
	if root == nil {
		return &Node{Key: key, Value: value, Parent: parent}
	} else if key < root.Key {
		root.Left = root.Left.insert(key, value, root)
	} else {
		root.Right = root.Right.insert(key, value, root)
	}
	return root
}

// Internal traversal method

func (root *Node) traverse(key int) *Node {
	if root == nil {
		panic("could not found the key") // do not know if a panic is the best solution for GO
	} else if key > root.Key {
		return root.Right.traverse(key)
	} else if key < root.Key {
		return root.Left.traverse(key)
	}

	return root
}

// Add add a new value to the tree
func (bst *BinarySearchTree) Add(key int, value string) {
	bst.Root = bst.Root.insert(key, value, nil)
}

// Search lookup for a Node with a specific key
func (bst *BinarySearchTree) Search(key int) *Node {
	return bst.Root.traverse(key)
}

func main() {
	var bst = BinarySearchTree{}
	bst.Add(6, "six")
	bst.Add(5, "five")
	bst.Add(7, "seven")
	bst.Add(3, "three")
	bst.Add(1, "one")

	var result = bst.Search(3)
	fmt.Println(result.Value)
	fmt.Println(result.isLeaf())
}
