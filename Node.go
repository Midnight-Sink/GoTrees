package GoTrees

// Node is a simple key-value struct used for the tree implementations.
type Node struct {
	Key         int
	Val         interface{}
	Left, Right *Node
}

// NewNode creates a new node with the key and value provided.
func NewNode(key int, val interface{}) Node {
	return Node{Key: key, Val: val}
}
