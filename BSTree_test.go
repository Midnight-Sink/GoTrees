package GoTrees

import (
	"math/rand"
	"sort"
	sc "strconv"
	"testing"
)

// NRAND is the number of random keys to test
const NRAND = 100

func TestBSTreeSlice(t *testing.T) {
	BST := NewBSTree()

	BST.Insert(10, nil)
	BST.Insert(11, nil)
	BST.Insert(9, nil)
	BST.Insert(8, nil)
	BST.Insert(14, nil)
	BST.Insert(12, nil)
	BST.Insert(13, nil)

	nodes := BST.Slice()
	expected := []int{8, 9, 10, 11, 12, 13, 14}
	for i, n := range nodes {
		if n == nil {
			t.Fatal("BST slice was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got nil node. ")
		} else if n.Key != expected[i] {
			t.Fatal("BST slice was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(n.Key) + ". ")
		}
	}
}

func TestBSTreeKeys(t *testing.T) {
	BST := NewBSTree()

	BST.Insert(10, nil)
	BST.Insert(11, nil)
	BST.Insert(9, nil)
	BST.Insert(8, nil)
	BST.Insert(14, nil)
	BST.Insert(12, nil)
	BST.Insert(13, nil)

	keys := BST.Keys()
	expected := []int{8, 9, 10, 11, 12, 13, 14}
	for i, k := range keys {
		if k != expected[i] {
			t.Fatal("BST key was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(k) + ". ")
		}
	}
}

func TestBSTreeValues(t *testing.T) {
	BST := NewBSTree()

	BST.Insert(10, 10)
	BST.Insert(11, 11)
	BST.Insert(9, 9)
	BST.Insert(8, 8)
	BST.Insert(14, 14)
	BST.Insert(12, 12)
	BST.Insert(13, 13)

	vals := BST.Values()
	expected := []int{8, 9, 10, 11, 12, 13, 14}
	for i, v := range vals {
		if v != expected[i] {
			t.Fatal("BST key was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(v.(int)) + ". ")
		}
	}
}

func TestBSTreeInsert(t *testing.T) {
	BST := NewBSTree()
	keys := make([]int, NRAND)

	for i := 0; i < NRAND; i++ {
		// NRAND - 1 to ensure at least one duplicate key
		key := rand.Intn(NRAND - 1)
		keys[i] = key
		BST.Insert(key, nil)
		if BST.Size != uint64(i+1) {
			t.Fatal("BST size incorrect, expected " + sc.Itoa(i+1) + " but got " + sc.Itoa(int(BST.Size)) + ". ")
		}
	}
	sort.Ints(keys)
	BSTkeys := BST.Keys()
	for i, k := range BSTkeys {
		if k != keys[i] {
			t.Fatal("BST key was incorrect, expected " + sc.Itoa(keys[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(k) + ". ")
		}
	}
}

func TestBSTreeFind(t *testing.T) {
	BST := NewBSTree()
	keys := make([]int, NRAND)

	for i := 0; i < NRAND; i++ {
		// NRAND - 1 to ensure at least one duplicate key
		key := rand.Intn(NRAND - 1)
		keys[i] = key
		BST.Insert(key, nil)
	}
	for _, k := range keys {
		if BST.Find(k) == nil {
			t.Fatal("Could not find node " + sc.Itoa(k) + ". ")
		}
	}
	if BST.Find(NRAND+1) != nil {
		t.Fatal("Found node that was not in the tree. ")
	}
}

func TestBSTreeContains(t *testing.T) {
	BST := NewBSTree()
	keys := make([]int, NRAND)

	for i := 0; i < NRAND; i++ {
		// NRAND - 1 to ensure at least one duplicate key
		key := rand.Intn(NRAND - 1)
		keys[i] = key
		BST.Insert(key, nil)
	}
	for _, k := range keys {
		if !BST.Contains(k) {
			t.Fatal("Could not find node " + sc.Itoa(k) + ". ")
		}
	}
	if BST.Contains(NRAND + 1) {
		t.Fatal("Found node that was not in the tree. ")
	}
}

func TestBSTreeDelete(t *testing.T) {
	BST := NewBSTree()
	keys := rand.Perm(NRAND)

	for _, key := range keys {
		BST.Insert(key, nil)
	}
	for _, k := range keys {
		if !BST.Delete(k) {
			t.Fatal("BST returned false when tree should have been modified. ")
		}
		// since this is permutation there are no duplicates
		if BST.Find(k) != nil {
			t.Fatal("Found deleted node after deletion of:" + sc.Itoa(k) + ". ")
		}
		/*
			if BST.Size != uint64(i+1) {
				t.Fatal("BST size incorrect, expected " + sc.Itoa(i+1) + " but got " + sc.Itoa(int(BST.Size)) + ". ")
			}
		*/
	}
}

func TestBSTreeDeleteAll(t *testing.T) {

}

func TestBSTreeString(t *testing.T) {

}