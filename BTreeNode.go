package GoTrees

type bTreeNode struct {
	nodes    []*node
	length   uint
	children []*bTreeNode
}

func NewbTreeNode(t int) bTreeNode {
	return bTreeNode{nodes: make([]*node, 2*t-1), length: 0, children: make([]*bTreeNode, 0)}
}

// Search will complete a binary search for the given key and return the node and its index in the list. If it is not found, it will return (nil, -1)
func (btNode *bTreeNode) Search(key int) (*node, int) {
	return nil, -1
}
