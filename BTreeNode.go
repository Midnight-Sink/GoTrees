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

// Search will complete a binary search for the given key and return the node and its index in the list. If it is not found, it will return (nil, -1)
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

	return nil, -1
}
