package GoTrees

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestBTreeNodeAdd(t *testing.T) {
	btn := newbTreeNode()
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
	btn := newbTreeNode()
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
	btn := newbTreeNode()
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
