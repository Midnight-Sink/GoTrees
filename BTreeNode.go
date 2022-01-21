package GoTrees

// bTreeNode is a container for the array of nodes used in b tree nodes
type bTreeNode struct {
	nodes       []*keyValue
	length      int
	children    []*bTreeNode
	numChildren int
}

func newbTreeNode() bTreeNode {
	return bTreeNode{nodes: make([]*keyValue, 0), length: 0, children: make([]*bTreeNode, 0), numChildren: 0}
}

// AddToList adds a node to the nodes list, it does not do any b-tree insert logic
func (btNode *bTreeNode) AddToList(n *keyValue) {
	min := 0
	max := btNode.length
	midPoint := (min + max) / 2

	for min < max {
		if btNode.nodes[midPoint].key > n.key {
			max = midPoint
		} else if btNode.nodes[midPoint].key < n.key {
			min = midPoint + 1
		} else {
			// this means node is duplicate key
			break
		}
		midPoint = (min + max) / 2
	}

	if btNode.length <= midPoint || btNode.length == 0 {
		btNode.nodes = append(btNode.nodes, n)
	} else {
		btNode.nodes = append(btNode.nodes[:midPoint+1], btNode.nodes[midPoint:]...)
		btNode.nodes[midPoint] = n
	}
	btNode.length++
}

// RemoveFromList removes an element from the nodes list, it does not do any b-tree delete logic
func (btNode *bTreeNode) RemoveFromList(key int) {
	_, i := btNode.Search(key)
	if i >= 0 {
		btNode.length--
		btNode.nodes = append(btNode.nodes[:i], btNode.nodes[i+1:]...)
	}
}

// Search will complete a binary search for the given key and return the node and its index in the list. If it is not found, it will return (nil, index of where the node would be)
func (btNode *bTreeNode) Search(key int) (*keyValue, int) {
	min := 0
	max := btNode.length
	midPoint := (min + max) / 2

	for min < max {
		if btNode.nodes[midPoint].key > key {
			max = midPoint
		} else if btNode.nodes[midPoint].key < key {
			min = midPoint + 1
		} else {
			return btNode.nodes[midPoint], midPoint
		}
		midPoint = (min + max) / 2
	}

	return nil, midPoint
}

func (btn *bTreeNode) SplitInTwo() (keyValue, *bTreeNode, *bTreeNode) {
	mid := int(btn.length / 2)
	var left *bTreeNode = nil
	var right *bTreeNode = nil
	if btn.numChildren > 0 {
		left = &bTreeNode{nodes: btn.nodes[:mid], length: btn.length / 2, children: btn.children[:mid], numChildren: btn.numChildren / 2}
		right = &bTreeNode{nodes: btn.nodes[mid+1:], length: btn.length / 2, children: btn.children[mid+1:], numChildren: btn.numChildren / 2}
	} else {
		left = &bTreeNode{nodes: btn.nodes[:mid], length: btn.length / 2, children: make([]*bTreeNode, 0), numChildren: 0}
		right = &bTreeNode{nodes: btn.nodes[mid+1:], length: btn.length / 2, children: make([]*bTreeNode, 0), numChildren: 0}
	}
	return *btn.nodes[mid], left, right
}

func (btn *bTreeNode) AddChild(other *bTreeNode) {
	btn.children = append(btn.children, other)
	btn.numChildren++
}

func (btn *bTreeNode) InsertTwoChildren(left *bTreeNode, right *bTreeNode, index int) {
	// making room for children nodes
	btn.children = append(btn.children, nil)
	for i := index; i < btn.numChildren; i++ {
		btn.children[i+1] = btn.children[i]
	}
	btn.children[index] = left
	btn.children[index+1] = right
	btn.numChildren += 1
}
