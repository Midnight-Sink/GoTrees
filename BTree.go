package GoTrees

// BTree is a b-tree using key-value nodes.
type BTree struct {
	Root []*Node
	Size uint64
	t    uint
}

// NewBTree returns an empty b-tree. t is the degree of the b-tree.
func NewBTree(t uint) BTree {
	return BTree{Root: make([]*Node, 0), Size: 0, t: t}
}

// Insert will insert node into the BT. If the node has a duplicate key, it will be placed on the RIGHT subtree.
func (bt *BTree) Insert(key int, value interface{}) {

}

// Delete will delete the closest occurance of the key to the root in the B-Tree. It will return whether or not the tree was changed.
func (bt *BTree) Delete(key int) bool {
	return false
}

// Find will find key in the B-Tree and return the node. Find will return the closest occurance of key to the root.
func (bt *BTree) Find(key int) *Node {
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

func (bt *BTree) Slice() []*Node {
	return nil
}

// Clear clears the B-Tree of all nodes.
func (bt *BTree) Clear() {
	bt.Root = nil
	bt.Size = 0
}

func (bt *BTree) Height() uint64 {
	return 0
}

// String will return the B-Tree represented as a string. Each level will be printed on a new line. Only keys will be printed. "X" represents a nil node. After the X is printed, all subsequent levels will not include this nodes children.
func (bt *BTree) String() string {
	return ""
}
