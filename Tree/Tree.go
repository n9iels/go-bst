package Tree

// Binary Search BinarySearchTree
type BinarySearchTree struct {
	Root *Node
}

// Add add a new value to the tree
func (bst *BinarySearchTree) Add(key int, value string) {
	bst.Root = bst.Root.insert(key, value, nil)
}

// Remove a node from the tree
func (bst *BinarySearchTree) Remove(key int) {
	bst.Root.remove(bst.Search(key))
}

// Search lookup for a Node with a specific key
func (bst *BinarySearchTree) Search(key int) *Node {
	return bst.Root.traverse(key)
}
