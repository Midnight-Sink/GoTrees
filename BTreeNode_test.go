package GoTrees

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestBTreeNodeAdd(t *testing.T) {
	btn := newbTreeNode(nAlloc * T)
	keys := rand.Perm(nRAND)

	for _, key := range keys {
		btn.AddToList(newKeyValue(key, key))
	}

	// This dummy value is ok since test values are positive
	last := -1
	for _, n := range btn.nodes {
		if n.key < last {
			t.Fatal("The nodes are not in order")
		}
		last = n.key
		n, _ := btn.Search(n.key)
		if n == nil {
			t.Fatal("A node is missing from the list")
		}
	}
}

func TestBTreeNodeRemove(t *testing.T) {
	btn := newbTreeNode(nAlloc * T)
	keys := rand.Perm(nRAND)

	btn.AddToList(newKeyValue(1, 1))
	btn.RemoveFromList(1)

	if btn.length != 0 {
		t.Fatal("After one addtion delete had the wrong length: " + strconv.Itoa(btn.length))
	}
	if n, _ := btn.Search(1); n != nil {
		t.Fatal("Search found a node after it was inserted and deleted: " + strconv.Itoa(btn.length))
	}

	for _, key := range keys {
		btn.AddToList(newKeyValue(key, key))
	}

	for _, n := range btn.nodes {
		btn.RemoveFromList(n.key)
		n, _ := btn.Search(n.key)
		if n != nil {
			t.Fatal("A node is still in the list after being deleted")
		}
	}
}

func TestBTreeNodeBinarySearch(t *testing.T) {
	btn := newbTreeNode(nAlloc * T)
	keys := rand.Perm(nRAND)

	for _, key := range keys {
		btn.AddToList(newKeyValue(key, key))
	}

	for _, key := range keys {
		// this test does not test the index correctly
		n, i := btn.Search(key)
		if n == nil {
			t.Fatal("Node search did not return the right result: node was nil")
		}
		if i == -1 || n.key != key {
			t.Fatal("Node search did not return the right result: returned key: " + strconv.Itoa(n.key) + " index: " + strconv.Itoa(i))
		}
	}

	if n, i := btn.Search(nRAND + 1); n != nil || i != -1 {
		t.Fatal("Node search did not return the right result, should have been (nil, -1)")
	}
}

func TestBTreeNodeSplitInTwo(t *testing.T) {
	btn := newbTreeNode(nAlloc * T)
	keys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, key := range keys {
		btn.AddToList(newKeyValue(key, key))
	}

	mid, left, right := btn.SplitInTwo(nAlloc * T)

	// mid
	if mid.key != 5 {
		t.Error("Mid key-value is not correct, expected 5 but got " + strconv.Itoa(mid.key))
	}
	// left
	realLen := len(left.nodes)
	if left.length != 4 {
		t.Error("Left numChildren incorrect, expected 4 but got " + strconv.Itoa(left.numChildren))
	} else if realLen != 4 {
		t.Error("Left children length incorrect, expected 4 but got " + strconv.Itoa(realLen))
	} else {
		for i, n := range left.nodes {
			if n.key != keys[i] {
				t.Error("Value not correct in left node: expected " + strconv.Itoa(keys[i]) + " but got " + strconv.Itoa(n.key))
			}
		}
	}
	// right
	realLen = len(right.nodes)
	if right.length != 4 {
		t.Error("Right numChildren incorrect, expected 4 but got " + strconv.Itoa(right.numChildren))
	} else if realLen != 4 {
		t.Error("Right children length incorrect, expected 4 but got " + strconv.Itoa(realLen))
	} else {
		for i, n := range right.nodes {
			if n.key != keys[i+5] {
				t.Error("Value not correct in right node: expected " + strconv.Itoa(keys[i+5]) + " but got " + strconv.Itoa(n.key))
			}
		}
	}
}
