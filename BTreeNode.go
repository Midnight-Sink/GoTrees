package GoTrees

// bTreeNode is a container for the array of nodes used in b tree nodes
type bTreeNode struct {
	nodes       []*keyValue
	length      int
	children    []*bTreeNode
	numChildren int
}

func newbTreeNode(alloc int) bTreeNode {
	return bTreeNode{nodes: make([]*keyValue, alloc), length: 0, children: make([]*bTreeNode, alloc), numChildren: 0}
}

// AddToList adds a node to the nodes list, it does not do any b-tree insert logic
func (btn *bTreeNode) AddToList(n *keyValue) {
	min := 0
	max := btn.length
	midPoint := (min + max) / 2

	for min < max {
		if btn.nodes[midPoint].key > n.key {
			max = midPoint
		} else if btn.nodes[midPoint].key < n.key {
			min = midPoint + 1
		} else {
			// this means node is duplicate key
			break
		}
		midPoint = (min + max) / 2
	}

	if btn.length <= midPoint || btn.length == 0 {
		btn.nodes = append(btn.nodes, n)
	} else {
		btn.nodes = append(btn.nodes[:midPoint+1], btn.nodes[midPoint:]...)
		btn.nodes[midPoint] = n
	}
	btn.length++
}

func (btn *bTreeNode) mergeRightSilbing(right *bTreeNode) {
	btn.nodes = append(btn.nodes, right.nodes...)
	btn.length += right.length
	btn.children = append(btn.children, right.children...)
	btn.numChildren += right.numChildren
}

// RemoveFromList removes an element from the nodes list, it does not do any b-tree delete logic
func (btn *bTreeNode) RemoveFromList(key int) {
	_, i := btn.Search(key)
	if i >= 0 {
		btn.length--
		btn.nodes = append(btn.nodes[:i], btn.nodes[i+1:]...)
	}
}

func (btn *bTreeNode) RemoveFromListAt(index int) {
	btn.length--
	btn.nodes = append(btn.nodes[:index], btn.nodes[index+1:]...)
}

func (btn *bTreeNode) ReplaceFromListAt(new *keyValue, index int) {
	btn.nodes[index] = new
}

// Search will complete a binary search for the given key and return the node and its index in the list. If it is not found, it will return (nil, index of where the node would be)
func (btn *bTreeNode) Search(key int) (*keyValue, int) {
	min := 0
	max := btn.length
	midPoint := (min + max) / 2

	for min < max {
		if btn.nodes[midPoint].key > key {
			max = midPoint
		} else if btn.nodes[midPoint].key < key {
			min = midPoint + 1
		} else {
			return btn.nodes[midPoint], midPoint
		}
		midPoint = (min + max) / 2
	}

	return nil, midPoint
}

// SplitInTwo splits a node into two subnodes, and takes the middle out
func (btn *bTreeNode) SplitInTwo(alloc int) (*keyValue, *bTreeNode, *bTreeNode) {
	mid := btn.length / 2
	var left *bTreeNode = nil
	var right *bTreeNode = nil

	nodeAlloc := max(alloc, btn.length/2)
	childAlloc := max(alloc, btn.length/2+1)

	if btn.numChildren > 0 {
		left = &bTreeNode{nodes: btn.nodes[:mid], length: btn.length / 2, children: btn.children[:mid+1], numChildren: btn.numChildren / 2}
		right = &bTreeNode{nodes: make([]*keyValue, nodeAlloc), length: btn.length / 2, children: make([]*bTreeNode, childAlloc), numChildren: btn.numChildren / 2}
		copy(right.children, btn.children[mid+1:])
	} else {
		// splitting a leaf, these nodes will need new child lists allocated
		left = &bTreeNode{nodes: btn.nodes[:mid], length: btn.length / 2, children: make([]*bTreeNode, btn.length/2+1), numChildren: 0}
		right = &bTreeNode{nodes: make([]*keyValue, nodeAlloc), length: btn.length / 2, children: make([]*bTreeNode, alloc+1), numChildren: 0}
	}
	// Only copy the right hand nodes, the left memory can be recycled
	copy(right.nodes, btn.nodes[mid+1:])
	return btn.nodes[mid], left, right
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// AddChild adds a child to the list
func (btn *bTreeNode) AddChild(other *bTreeNode) {
	btn.children = append(btn.children, other)
	btn.numChildren++
}

// DeleteChild removes a child at index from the child list
func (btn *bTreeNode) DeleteChild(index int) {
	if index >= 0 {
		btn.numChildren--
		btn.children = append(btn.children[:index], btn.children[index+1:]...)
	}
}

// InsertTwoChildren adds 2 children to the child list, overwriting the child at index
func (btn *bTreeNode) InsertTwoChildren(left *bTreeNode, right *bTreeNode, index int) {
	// making room for children nodes
	btn.children = append(btn.children, nil)
	// shift over current children
	for i := btn.numChildren - 1; i >= index; i-- {
		btn.children[i+1] = btn.children[i]
	}
	// assign new children (note this will overwrite an existing node on purpose)
	btn.children[index] = left
	btn.children[index+1] = right
	btn.numChildren += 1
}
