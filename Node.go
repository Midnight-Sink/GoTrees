package GoTrees

// node is a simple key-value struct used for the tree implementations.
type node struct {
	Key         int
	Val         interface{}
	Left, Right *node
}

// newNode creates a new node with the key and value provided.
func newNode(key int, val interface{}) *node {
	return &node{Key: key, Val: val}
}
