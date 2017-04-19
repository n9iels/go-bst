package Tree

type Node struct {
	Key                 int
	Value               string
	Left, Right, Parent *Node
}

// Insert a new node
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

// Remove a node
func (root *Node) remove(node *Node) {
	if node.IsLeaf() {
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	} else if node.Left == nil {
		node.Right.Parent = node.Parent

		if node.Parent.Left == node {
			node.Parent.Left = node.Right
		} else {
			node.Parent.Right = node.Right
		}
	} else if node.Right == nil {
		node.Left.Parent = node.Parent

		if node.Parent.Right == node {
			node.Parent.Right = node.Left
		} else {
			node.Parent.Left = node.Left
		}
	} else {
		var successor *Node = node.Left

		for successor.Right != nil {
			successor = successor.Right
		}

		successor.Left = node.Left
		successor.Right = node.Right
		successor.Parent = node.Parent

		if node.Parent.Left == node {
			node.Parent.Left = successor
		} else {
			node.Parent.Right = successor
		}
	}
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

// External method to determine if a node is a leaf or not
func (root *Node) IsLeaf() bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
