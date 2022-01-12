package GoTrees

import "strconv"

// BTree is a b-tree using key-value nodes.
type BTree struct {
	root *bTreeNode
	size uint64
	t    uint
}

// NewBTree returns an empty b-tree. t is the degree of the b-tree.
func NewBTree(t uint) BTree {
	return BTree{root: nil, size: 0, t: t}
}

// Insert will insert node into the BT. If the node has a duplicate key, it will be placed on the RIGHT subtree.
func (bt *BTree) Insert(key int, value interface{}) {
	bt.size++
}

// Delete will delete the closest occurance of the key to the root in the B-Tree. It will return whether or not the tree was changed.
func (bt *BTree) Delete(key int) bool {
	bt.size--
	return false
}

// Find will find key in the B-Tree and return the node. Find will return the closest occurance of key to the root.
func (bt *BTree) Find(key int) *node {
	return nil
}

// Contains determines if key exists in the B-Tree and returns the result.
func (bt *BTree) Contains(key int) bool {
	return bt.Find(key) != nil
}

func (bt *BTree) Keys() []int {
	return nil
}

func (bt *BTree) Values() []interface{} {
	return nil
}

func (bt *BTree) slice() []*keyValue {
	nodes := make([]*keyValue, bt.size)
	i := 0
	nodeStack := []*bTreeNode{}
	indexStack := []int{}
	stacksize := 0
	n := bt.root
	for n != nil || stacksize != 0 {
		if n != nil {
			for _, child := range n.children {
				nodeStack = append(nodeStack, child)
				indexStack = append(indexStack, 0)
				stacksize++
			}
			n = n.Left
		} else {
			n = nodeStack[stacksize-1]
			nodeStack = nodeStack[:stacksize-1]
			stacksize--
			nodes[i] = n
			i++
			n = n.Right
		}
	}
	return nodes
}

// Clear clears the B-Tree of all nodes.
func (bt *BTree) Clear() {
	bt.root = nil
	bt.size = 0
}

func (bt *BTree) Height() uint64 {
	if bt.root == nil {
		return 0
	}

	height := uint64(0)
	nodeQ := []*bTreeNode{}
	qsize := 0

	nodeQ = append(nodeQ, bt.root)
	qsize++

	for {
		if qsize == 0 {
			return height
		}
		nodeCount := qsize
		for nodeCount > 0 {
			for _, child := range nodeQ[0].children {
				nodeQ = append(nodeQ, child)
				qsize++
			}
			nodeQ = nodeQ[1:]
			qsize--
			nodeCount--
		}
	}
}

// String will return the B-Tree represented as a string. Each level will be printed on a new line. Only keys will be printed. "X" represents a nil node. After the X is printed, all subsequent levels will not include this nodes children.
func (bt *BTree) String() string {
	nodeQ := []*bTreeNode{}
	qsize := 0
	str := ""

	nodeQ = append(nodeQ, bt.root)
	qsize++

	for {
		if qsize == 0 {
			return str
		}
		nodeCount := qsize
		for nodeCount > 0 {
			if nodeQ[0] == nil {
				str = str + "X "
			} else {
				for _, child := range nodeQ[0].children {
					nodeQ = append(nodeQ, child)
					qsize++
				}
				str += "["
				for _, kv := range nodeQ[0].nodes {
					str = str + strconv.Itoa(kv.key) + " "
				}
				str += "] "
			}
			nodeQ = nodeQ[1:]
			qsize--
			nodeCount--
		}
		str = str + "\n"
	}
}

func (bt *BTree) Size() uint64 {
	return bt.size
}
