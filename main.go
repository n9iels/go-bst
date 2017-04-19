package main

import (
	"fmt"
	"github.com/n9iels/go-bst/Tree"
)

// Node for a BinarySearchTree

func main() {
	var bst = Tree.BinarySearchTree{}
	bst.Add(10, "then")
	bst.Add(7, "seven")
	bst.Add(8, "eight")
	bst.Add(5, "five")
	bst.Add(4, "four")
	//bst.Add(6, "six")
	bst.Add(11, "eleven")
	bst.Add(12, "twelf")
	bst.Add(13, "thirdtheen")
	bst.Remove(7)

	var result = bst.Search(12)
	fmt.Println(result.Value)
	fmt.Println(result.IsLeaf())
}
