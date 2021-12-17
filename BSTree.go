package GoTrees

// BSTree is a binary search tree using key-value nodes.
type BSTree struct {
	Root *Node
	Size uint64
}

// NewBSTree returns an empty binary search tree. The values will be initialized the same way when doing BSTree{}
func NewBSTree() BSTree {
	return BSTree{Root: nil, Size: 0}
}

// Insert will insert node into the BST. If the node has a duplicate key, it will be placed on the RIGHT subtree.
func (bst *BSTree) Insert(key int, value interface{}) {
	node := NewNode(key, value)
	bst.Size++
	// n is the iterating node variable
	n := bst.Root
	for n != nil {
		// check which side the node should progress to
		if n.Key <= node.Key {
			// check if the node can be added to the right side
			if n.Right == nil {
				n.Right = &node
				return
			} else {
				n = n.Right
			}
		} else {
			// check if the node can be added to the left side
			if n.Left == nil {
				n.Left = &node
				return
			} else {
				n = n.Left
			}
		}
	}
	// if the tree is empty
	bst.Root = &node
}

// Delete will delete the closest occurance of the key to the root in the BST. It will return whether or not the tree was changed.
func (bst *BSTree) Delete(key int) bool {
	// the side of the child in relation to the parent (false is left)
	side := false
	// the parent node for the current node
	var parent *Node = nil
	// the current node
	curr := bst.Root
	for curr != nil {
		// check which side the node should progress to
		if curr.Key <= key {
			// check if the node has the desired key
			if curr.Key == key {
				bst.Size--
				// find in order successor
				var ios *Node = curr.Right
				if ios != nil {
					var parent_ios *Node = curr
					// in order successor is leftmost node in subtree
					for ios.Left != nil {
						parent_ios = ios
						ios = ios.Left
					}
					parent_ios.Left = ios.Right
				}
				ios.Left = curr.Left
				ios.Right = curr.Right
				if parent == nil {
					bst.Root = ios
				} else if side {
					parent.Right = ios
				} else {
					parent.Left = ios
				}
				return true
			} else if curr.Right == nil {
				return false
			} else {
				side = true
				parent = curr
				curr = curr.Right
			}
		} else {
			if curr.Left == nil {
				return false
			} else {
				side = false
				parent = curr
				curr = curr.Left
			}
		}
	}
	// BST is empty
	return false
}

// DeleteAll will delete all occurances of the key in the BST. It will return whether or not the tree was changed.
func (bst *BSTree) DeleteAll(key int) bool {
	return false
}

// Find will find key in the BST and return the node. Find will return the closest occurance of key to the root.
func (bst *BSTree) Find(key int) *Node {
	// n is the iterating node variable
	n := bst.Root
	for n != nil {
		// check which side the node should progress to
		if n.Key <= key {
			// check if the node has the desired key
			if n.Key == key {
				return n
			} else if n.Right == nil {
				return nil
			} else {
				n = n.Right
			}
		} else {
			if n.Left == nil {
				return nil
			} else {
				n = n.Left
			}
		}
	}
	// if the tree is empty
	return nil
}

// Contains determines if key exists in the BST and returns the result.
func (bst *BSTree) Contains(key int) bool {
	return bst.Find(key) != nil
}

func (bst *BSTree) Keys() []int {
	keys := make([]int, bst.Size)
	i := 0
	nodeStack := []*Node{}
	stackSize := 0
	n := bst.Root
	for n != nil || stackSize != 0 {
		if n != nil {
			nodeStack = append(nodeStack, n)
			stackSize++
			n = n.Left
		} else {
			n = nodeStack[stackSize-1]
			nodeStack = nodeStack[:stackSize-1]
			stackSize--
			keys[i] = n.Key
			i++
			n = n.Right
		}
	}
	return keys
}

func (bst *BSTree) Values() []interface{} {
	vals := make([]interface{}, bst.Size)
	i := 0
	nodeStack := []*Node{}
	stackSize := 0
	n := bst.Root
	for n != nil || stackSize != 0 {
		if n != nil {
			nodeStack = append(nodeStack, n)
			stackSize++
			n = n.Left
		} else {
			n = nodeStack[stackSize-1]
			nodeStack = nodeStack[:stackSize-1]
			stackSize--
			vals[i] = n.Val
			i++
			n = n.Right
		}
	}
	return vals
}

func (bst *BSTree) Slice() []*Node {
	nodes := make([]*Node, bst.Size)
	i := 0
	nodeStack := []*Node{}
	stackSize := 0
	n := bst.Root
	for n != nil || stackSize != 0 {
		if n != nil {
			nodeStack = append(nodeStack, n)
			stackSize++
			n = n.Left
		} else {
			n = nodeStack[stackSize-1]
			nodeStack = nodeStack[:stackSize-1]
			stackSize--
			nodes[i] = n
			i++
			n = n.Right
		}
	}
	return nodes
}

// Clear clears the BST of all nodes.
func (bst *BSTree) Clear() {
	bst.Root = nil
	bst.Size = 0
}

func (bst *BSTree) Height() {

}

// String will return the BST represented as a string. Each level will be printed on a new line. Only keys will be printed. "X" represents a nil node.
func (bst *BSTree) String() {

}

// getLevel returns all nodes on a level of the tree. If an element is nil in the node slice it means there is no node there
func (bst *BSTree) getLevel(level uint64, n *Node) []*Node {
	if n == nil {
		return []*Node{nil}
	} else if level == 0 {
		return []*Node{n}
	}
	l := bst.getLevel(level-1, n.Left)
	r := bst.getLevel(level-1, n.Right)
	return append(append(l, n), r...)
}
