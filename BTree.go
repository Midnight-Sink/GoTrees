package GoTrees

import (
	"strconv"
)

// BTree is a b-tree using key-value nodes.
type BTree struct {
	root *bTreeNode
	size uint64
	t    uint
}

// NewBTree returns an empty b-tree. The degree of the b tree is 2*t+2. (This ensures valid max-degree. Since this b-tree splits preemptively the degree must be even so it will split with an odd number of pairs)
func NewBTree(t uint) BTree {
	root := newbTreeNode()
	return BTree{root: &root, size: 0, t: 2*t + 2}
}

// Insert will insert node into the BT. If the node has a duplicate key, it will be placed on the RIGHT subtree.
func (bt *BTree) Insert(key int, value interface{}) {
	// check the root for capacity (a new node will be allocated)
	if bt.root.length > int(bt.t) {
		mid, left, right := bt.root.SplitInTwo()
		newRoot := newbTreeNode()
		bt.root = &newRoot
		bt.root.AddToList(&mid)
		bt.root.AddChild(left)
		bt.root.AddChild(right)
	}
	curr := bt.root
	for curr.numChildren != 0 {
		_, indexNext := curr.Search(key)
		if curr.children[indexNext].length > int(bt.t) {
			// split the node
			mid, left, right := curr.children[indexNext].SplitInTwo()
			curr.AddToList(&mid)
			curr.InsertTwoChildren(left, right, indexNext)
			// determine which new node is the next child
			if mid.key <= key {
				curr = right
			} else {
				curr = left
			}
		} else {
			// progress to the next child node
			curr = curr.children[indexNext]
		}
	}
	bt.size++
	// since this B tree preemtively splits nodes, this key-value will fit into this node
	curr.AddToList(newKeyValue(key, value))
}

// Delete will delete the closest occurance of the key to the root in the B-Tree. It will return whether or not the tree was changed.
func (bt *BTree) Delete(key int) bool {
	bt.size--
	return false
}

// Find will find key in the B-Tree and return the node. Find will return the closest occurance of key to the root.
func (bt *BTree) Find(key int) *interface{} {
	curr := bt.root

	for {
		res, i := curr.Search(key)
		if res != nil {
			// the node was found
			return &res.value
		} else {
			if curr.numChildren == 0 {
				// the node wasn't found and there are no more children to check
				return nil
			} else {
				// check the next child
				curr = curr.children[i]
			}
		}
	}
}

// Contains determines if key exists in the B-Tree and returns the result.
func (bt *BTree) Contains(key int) bool {
	return bt.Find(key) != nil
}

func (bt *BTree) Keys() []int {
	keys := make([]int, bt.size)
	if bt.size == 0 {
		return keys
	}
	i := 0
	nodeStack := []*bTreeNode{}
	indexStack := []int{}
	stacksize := 0

	nodeStack = append(nodeStack, bt.root)
	indexStack = append(indexStack, 0)
	stacksize++

	for stacksize > 0 {
		curr := nodeStack[stacksize-1]
		if curr.length < indexStack[stacksize-1] {
			// this node has had all of its children accounted for and can be popped
			nodeStack = nodeStack[:stacksize-1]
			indexStack = indexStack[:stacksize-1]
			stacksize--
		} else {
			// this node may have children to be accounted for
			if indexStack[stacksize-1] > 0 {
				// parent node values inbetween the children
				keys[i] = curr.nodes[indexStack[stacksize-1]-1].key
				i++
			}
			if curr.numChildren > 0 {
				// increment the parent index value
				indexStack[stacksize-1]++
				// this node is inner (or root) and must check children
				nodeStack = append(nodeStack, curr.children[indexStack[stacksize-1]])
				indexStack = append(indexStack, 0)
				stacksize++
			} else {
				// this node is a leaf and can have all of the kv pairs added
				for _, kv := range curr.nodes {
					keys[i] = kv.key
					i++
				}
				nodeStack = nodeStack[:stacksize-1]
				indexStack = indexStack[:stacksize-1]
				stacksize--
			}
		}
	}

	return keys
}

func (bt *BTree) Values() []interface{} {
	vals := make([]interface{}, bt.size)
	if bt.size == 0 {
		return vals
	}
	i := 0
	nodeStack := []*bTreeNode{}
	indexStack := []int{}
	stacksize := 0

	nodeStack = append(nodeStack, bt.root)
	indexStack = append(indexStack, 0)
	stacksize++

	for stacksize > 0 {
		curr := nodeStack[stacksize-1]
		if curr.length < indexStack[stacksize-1] {
			// this node has had all of its children accounted for and can be popped
			nodeStack = nodeStack[:stacksize-1]
			indexStack = indexStack[:stacksize-1]
			stacksize--
		} else {
			// this node may have children to be accounted for
			if indexStack[stacksize-1] > 0 {
				// parent node values inbetween the children
				vals[i] = curr.nodes[indexStack[stacksize-1]-1].value
				i++
			}
			if curr.numChildren > 0 {
				// increment the parent index value
				indexStack[stacksize-1]++
				// this node is inner (or root) and must check children
				nodeStack = append(nodeStack, curr.children[indexStack[stacksize-1]])
				indexStack = append(indexStack, 0)
				stacksize++
			} else {
				// this node is a leaf and can have all of the kv pairs added
				for _, kv := range curr.nodes {
					vals[i] = kv.value
					i++
				}
				nodeStack = nodeStack[:stacksize-1]
				indexStack = indexStack[:stacksize-1]
				stacksize--
			}
		}
	}

	return vals
}

func (bt *BTree) slice() []*keyValue {
	nodes := make([]*keyValue, bt.size)
	if bt.size == 0 {
		return nodes
	}
	i := 0
	nodeStack := []*bTreeNode{}
	indexStack := []int{}
	stacksize := 0

	nodeStack = append(nodeStack, bt.root)
	indexStack = append(indexStack, 0)
	stacksize++

	for stacksize > 0 {
		curr := nodeStack[stacksize-1]
		if curr.length < indexStack[stacksize-1] {
			// this node has had all of its children accounted for and can be popped
			nodeStack = nodeStack[:stacksize-1]
			indexStack = indexStack[:stacksize-1]
			stacksize--
		} else {
			// this node may have children to be accounted for
			if indexStack[stacksize-1] > 0 {
				// parent node values inbetween the children
				nodes[i] = curr.nodes[indexStack[stacksize-1]-1]
				i++
			}
			if curr.numChildren > 0 {
				// increment the parent index value
				indexStack[stacksize-1]++
				// this node is inner (or root) and must check children
				nodeStack = append(nodeStack, curr.children[indexStack[stacksize-1]])
				indexStack = append(indexStack, 0)
				stacksize++
			} else {
				// this node is a leaf and can have all of the kv pairs added
				for _, kv := range curr.nodes {
					nodes[i] = kv
					i++
				}
				nodeStack = nodeStack[:stacksize-1]
				indexStack = indexStack[:stacksize-1]
				stacksize--
			}
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
