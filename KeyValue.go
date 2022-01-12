package GoTrees

type keyValue struct {
	key   int
	value interface{}
}

func newKeyValue(key int, value interface{}) *keyValue {
	return &keyValue{key: key, value: value}
}
