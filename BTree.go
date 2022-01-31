package GoTrees

import (
	"strconv"
)

// BTree is a b-tree using key-value nodes.
type BTree struct {
	root      *bTreeNode
	size      uint64
	t         uint
	initAlloc int
}

// NewBTree returns an empty b-tree. The degree of the b tree is 2*t+2. (This ensures valid max-degree. Since this b-tree splits preemptively the degree must be even so it will split with an odd number of pairs)
func NewBTree(t uint, alloc float32) BTree {
	if alloc > 1 {
		alloc = 1
	} else if alloc < 0 {
		alloc = 0
	}
	root := newbTreeNode(int(alloc * float32(t)))
	return BTree{root: &root, size: 0, t: 2*t + 2, initAlloc: int(alloc * float32(t))}
}

// Insert will insert node into the BT. A duplicate tree could be placed in the left or right subtree to maintain balance.
func (bt *BTree) Insert(key int, value interface{}) {
	// check the root for capacity (a new node will be allocated)
	if bt.root.length > int(bt.t) {
		mid, left, right := bt.root.SplitInTwo(bt.initAlloc)
		newRoot := newbTreeNode(bt.initAlloc)
		bt.root = &newRoot
		bt.root.AddToList(mid)
		bt.root.AddChild(left)
		bt.root.AddChild(right)
	}
	curr := bt.root
	for curr.numChildren != 0 {
		_, indexNext := curr.Search(key)
		if curr.children[indexNext].length > int(bt.t) {
			// split the node
			mid, left, right := curr.children[indexNext].SplitInTwo(bt.initAlloc)
			curr.AddToList(mid)
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
				// this node is inner (or root) and must check children
				nodeStack = append(nodeStack, curr.children[indexStack[stacksize-1]])
				indexStack = append(indexStack, 0)
				// increment the parent index value
				indexStack[stacksize-1]++
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
				// this node is inner (or root) and must check children
				nodeStack = append(nodeStack, curr.children[indexStack[stacksize-1]])
				indexStack = append(indexStack, 0)
				// increment the parent index value
				indexStack[stacksize-1]++
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
				// this node is inner (or root) and must check children
				nodeStack = append(nodeStack, curr.children[indexStack[stacksize-1]])
				indexStack = append(indexStack, 0)
				// increment the parent index value
				indexStack[stacksize-1]++
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

// Height calculates the height of the B tree
func (bt *BTree) Height() uint64 {
	if bt.root == nil || bt.root.length == 0 {
		return 0
	}
	height := uint64(0)
	curr := bt.root

	for curr.numChildren > 0 {
		height++
		// height calcuated by following left side of tree (works since tree is always pefectly balanced)
		curr = curr.children[0]
	}
	// Adding 1 for the missed iteration on the leaf node, and adding 1 more since the loop counts "links" rather than nodes
	return height + 2
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

// Delete will delete the closest occurance of the key to the root in the B-Tree. It will return whether or not the tree was changed.
func (bt *BTree) Delete(key int) bool {
	curr := bt.root
	t := int(bt.t)

	for {
		res, i := curr.Search(key)
		leftSibling := i > 0
		rightSibling := i < curr.length
		if res != nil {
			// the node was found
			if curr.numChildren == 0 {
				// this node is a leaf node. since this premtively merges nodes, there will be room for deletion
				curr.RemoveFromListAt(i)
			} else {
				// this node is an interior node, will replace with another node. The orientation is as follows:
				//			[#] <-- curr traversal is here
				// 			/ \
				//		  [#] [#] <-- looking at these children
				if leftSibling && curr.children[i-1].length > t {
					// replace the deleted node with the in order predecessor
					pred := findAndDeleteIOP(curr.children[i-1], t)
					curr.ReplaceFromListAt(pred, i)
				} else if rightSibling && curr.children[i+1].length > t {
					// replace the deleted node with the in order successor
					succ := findAndDeleteIOS(curr.children[i+1], t)
					curr.ReplaceFromListAt(succ, i)
				} else {
					// merge children and push this KV down 1 level since neither sibling can fill the gap
					if leftSibling {
						parentMerge(curr, curr.children[i-1], curr.children[i], i)
					} else {
						parentMerge(curr, curr.children[i], curr.children[i+1], i)
					}
					// must skip return since more merges may be required
					curr = curr.children[i]
					continue
				}
			}
			bt.size--
			return true
		} else {
			if curr.numChildren == 0 {
				// the node wasn't found and there are no more children to check
				return false
			} else {
				// check the next child size
				// The orientation is as follows:
				// 			[#] <-- curr is here
				//		   / | \
				//	   [#?] [#] [#?] <-- traversal is going to the middle child, looking at the two (possible) siblings
				if curr.children[i].length <= t {
					// premptive merging is required
					validateNextChildSize(curr, leftSibling, rightSibling, i, t)
				}
				curr = curr.children[i]
			}
		}
	}
}

func validateNextChildSize(curr *bTreeNode, leftSibling, rightSibling bool, i int, t int) {
	if curr.children[i].length <= t {
		// premptive merging is required
		if leftSibling && curr.children[i-1].length > t {
			// there is a left sibling with capacity
			borrowLeft(curr, curr.children[i-1], curr.children[i], i)
		} else if rightSibling && curr.children[i+1].length > t {
			// there is a right sibling with capacity
			borrowRight(curr, curr.children[i+1], curr.children[i], i)
		} else {
			// must merge with one sibling
			if leftSibling {
				parentMerge(curr, curr.children[i-1], curr.children[i], i)
			} else {
				parentMerge(curr, curr.children[i], curr.children[i+1], i)
			}
		}
	}
}

// borrowLeft is called when the current node can borrow a KV from the parent who can then borrow a KV from the left sibling of curr
func borrowLeft(parent, left, curr *bTreeNode, index int) {
	// shift the in order predecessor down to this current node
	curr.AddToList(parent.nodes[index])
	// replace the shifted parent KV with the in order predecessor (largest key in left)
	parent.nodes[index] = left.nodes[left.length-1]
	// remove the KV from left
	left.RemoveFromListAt(left.length - 1)
}

// borrowRight is called when the current node can borrow a KV from the parent who can then borrow a KV from the right sibling of curr
func borrowRight(parent, right, curr *bTreeNode, index int) {
	// shift the in order successor down to this current node
	curr.AddToList(parent.nodes[index])
	// replace the shifted parent KV with the in order successor (smallest key in right)
	parent.nodes[index] = right.nodes[0]
	// remove the KV from right
	right.RemoveFromListAt(0)
}

// parentMerge is called when it cannot borrow from both left and right silbings
func parentMerge(parent, left, curr *bTreeNode, index int) {
	// add the node from the parent in the merge
	left.AddToList(parent.nodes[index])
	parent.RemoveFromListAt(index)
	// merge the silbing to the right (curr)
	left.mergeRightSilbing(curr)
	// delete curr as it has been merged into left
	parent.DeleteChild(index)
}

// findAndDeleteIOP find and delete in order predecessor
func findAndDeleteIOP(start *bTreeNode, t int) *keyValue {
	for start.numChildren != 0 {
		validateNextChildSize(start, true, false, start.numChildren-1, t)
		start = start.children[start.numChildren-1]
	}
	pred := start.nodes[start.length-1]
	start.RemoveFromListAt(start.length - 1)
	return pred
}

// findAndDeleteIOS find and delete in order successor
func findAndDeleteIOS(start *bTreeNode, t int) *keyValue {
	for start.numChildren != 0 {
		validateNextChildSize(start, false, true, 0, t)
		start = start.children[0]
	}
	pred := start.nodes[0]
	start.RemoveFromListAt(0)
	return pred
}
