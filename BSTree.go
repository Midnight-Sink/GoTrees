package GoTrees

import (
	"strconv"
)

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
					if parent_ios != curr {
						parent_ios.Left = ios.Right
					} else {
						parent_ios.Right = ios.Right
					}

					ios.Left = curr.Left
					// circular reference here
					ios.Right = curr.Right
					if parent == nil {
						bst.Root = ios
					} else if side {
						parent.Right = ios
					} else {
						parent.Left = ios
					}
				} else {
					if parent == nil {
						bst.Root = curr.Left
					} else if side {
						parent.Right = curr.Left
					} else {
						parent.Left = curr.Left
					}
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

func (bst *BSTree) Height() uint64 {
	if bst.Root == nil {
		return 0
	}
	height := uint64(0)
	nodeQ := []*Node{}
	qSize := 0

	nodeQ = append(nodeQ, bst.Root)
	qSize++

	for {
		if qSize == 0 {
			return height
		}
		height++
		nodeCount := qSize
		for nodeCount > 0 {
			if nodeQ[0].Left != nil {
				nodeQ = append(nodeQ, nodeQ[0].Left)
				qSize++
			}
			if nodeQ[0].Right != nil {
				nodeQ = append(nodeQ, nodeQ[0].Right)
				qSize++
			}
			nodeQ = nodeQ[1:]
			qSize--
			nodeCount--
		}
	}
}

// String will return the BST represented as a string. Each level will be printed on a new line. Only keys will be printed. "X" represents a nil node. After the X is printed, all subsequent levels will not include this nodes children.
func (bst *BSTree) String() string {
	nodeQ := []*Node{}
	qSize := 0
	str := ""

	nodeQ = append(nodeQ, bst.Root)
	qSize++

	for {
		if qSize == 0 {
			return str
		}
		nodeCount := qSize
		for nodeCount > 0 {
			if nodeQ[0] == nil {
				str = str + "X "
			} else {
				nodeQ = append(nodeQ, nodeQ[0].Left)
				qSize++
				nodeQ = append(nodeQ, nodeQ[0].Right)
				qSize++
				str = str + strconv.Itoa(nodeQ[0].Key) + " "
			}
			nodeQ = nodeQ[1:]
			qSize--
			nodeCount--
		}
		str = str + "\n"
	}
}
