package GoTrees

import (
	"math/rand"
	"sort"
	"strconv"
	sc "strconv"
	"testing"
)

const T = 0
const nAlloc = .5

func TestBTreeSlice(t *testing.T) {
	BT := NewBTree(T, nAlloc)

	BT.Insert(10, nil)
	BT.Insert(11, nil)
	BT.Insert(9, nil)
	BT.Insert(8, nil)
	BT.Insert(14, nil)
	BT.Insert(12, nil)
	BT.Insert(13, nil)

	nodes := BT.slice()
	expected := []int{8, 9, 10, 11, 12, 13, 14}
	for i, n := range nodes {
		if n == nil {
			t.Fatal("BT slice was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got nil node. ")
		} else if n.key != expected[i] {
			t.Fatal("BT slice was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(n.key) + ". ")
		}
	}
}

func TestBTreeKeys(t *testing.T) {
	BT := NewBTree(T, nAlloc)

	BT.Insert(10, nil)
	BT.Insert(11, nil)
	BT.Insert(9, nil)
	BT.Insert(8, nil)
	BT.Insert(14, nil)
	BT.Insert(12, nil)
	BT.Insert(13, nil)

	keys := BT.Keys()
	expected := []int{8, 9, 10, 11, 12, 13, 14}
	for i, k := range keys {
		if k != expected[i] {
			t.Fatal("BT key was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(k) + ". ")
		}
	}
}

func TestBTreeValues(t *testing.T) {
	BT := NewBTree(T, nAlloc)

	BT.Insert(10, 10)
	BT.Insert(11, 11)
	BT.Insert(9, 9)
	BT.Insert(8, 8)
	BT.Insert(14, 14)
	BT.Insert(12, 12)
	BT.Insert(13, 13)

	vals := BT.Values()
	expected := []int{8, 9, 10, 11, 12, 13, 14}
	for i, v := range vals {
		if v != expected[i] {
			t.Fatal("BT key was incorrect, expected " + sc.Itoa(expected[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(v.(int)) + ". ")
		}
	}
}

func TestBTreeInsert(t *testing.T) {
	BT := NewBTree(T, nAlloc)
	keys := make([]int, nRAND)

	for i := 0; i < nRAND; i++ {
		// nRAND - 1 to ensure at least one duplicate key
		key := rand.Intn(nRAND - 1)
		keys[i] = key
		BT.Insert(key, nil)
		if BT.size != uint64(i+1) {
			t.Fatal("BT size incorrect, expected " + sc.Itoa(i+1) + " but got " + sc.Itoa(int(BT.size)) + ". ")
		}
		// t.Log("Inserted key: " + strconv.Itoa(key))
		// t.Log(BT.String())
		// t.Log("==================================")
	}
	sort.Ints(keys)
	BTkeys := BT.Keys()
	for i, k := range BTkeys {
		if k != keys[i] {
			t.Fatal("BT key was incorrect, expected " + sc.Itoa(keys[i]) + " at index " + sc.Itoa(i) + " but got " + sc.Itoa(k) + ". ")
		}
	}
}

func TestBTreeHeight(t *testing.T) {
	BT := NewBTree(T, nAlloc)

	h := BT.Height()
	if h != 0 {
		t.Fatal("Height was expected to be 0 but was " + sc.Itoa(int(h)))
	}

	keys := make([]int, nRAND)
	for i := 0; i < nRAND; i++ {
		// nRAND - 1 to ensure at least one duplicate key
		key := rand.Intn(nRAND - 1)
		keys[i] = key
		BT.Insert(key, nil)
		if BT.size != uint64(i+1) {
			t.Fatal("BT size incorrect, expected " + sc.Itoa(i+1) + " but got " + sc.Itoa(int(BT.size)) + ". ")
		}
	}

	h = BT.Height()
	if h != 6 {
		t.Fatal("Height was expected to be 6 but was " + sc.Itoa(int(h)))
	}
}

func TestBTreeFind(t *testing.T) {
	BT := NewBTree(T, nAlloc)
	keys := make([]int, nRAND)

	for i := 0; i < nRAND; i++ {
		// nRAND - 1 to ensure at least one duplicate key
		key := rand.Intn(nRAND - 1)
		keys[i] = key
		BT.Insert(key, nil)
	}
	for _, k := range keys {
		if BT.Find(k) == nil {
			t.Fatal("Could not find node " + sc.Itoa(k) + ". ")
		}
	}
	if BT.Find(nRAND+1) != nil {
		t.Fatal("Found node that was not in the tree. ")
	}
}

func TestBTreeContains(t *testing.T) {
	BT := NewBTree(T, nAlloc)
	keys := make([]int, nRAND)

	for i := 0; i < nRAND; i++ {
		// nRAND - 1 to ensure at least one duplicate key
		key := rand.Intn(nRAND - 1)
		keys[i] = key
		BT.Insert(key, nil)
	}
	for _, k := range keys {
		if !BT.Contains(k) {
			t.Fatal("Could not find node " + sc.Itoa(k) + ". ")
		}
	}
	if BT.Contains(nRAND + 1) {
		t.Fatal("Found node that was not in the tree. ")
	}
}

func TestBTreeDelete(t *testing.T) {
	BT := NewBTree(T, nAlloc)
	keys := rand.Perm(nRAND)

	for _, key := range keys {
		BT.Insert(key, nil)
	}
	for i, k := range keys {
		if !BT.Delete(k) {
			t.Fatal("BT returned false when tree should have been modified. ")
		}
		// since this is permutation there are no duplicates
		if BT.Find(k) != nil {
			t.Fatal("Found deleted node after deletion of:" + sc.Itoa(k) + ". ")
		}
		if BT.size != uint64(nRAND-(i+1)) {
			t.Fatal("BT size incorrect, expected " + sc.Itoa(nRAND-(i+1)) + " but got " + sc.Itoa(int(BT.size)) + ". ")
		}
		for _, kInner := range keys[i+1:] {
			if !BT.Contains(kInner) {
				t.Fatal("Tree is missing key that wasn't deleted yet: " + strconv.Itoa(kInner))
			}
		}
	}
}

func TestBTreeString(t *testing.T) {
	BT := NewBTree(T, nAlloc)

	BT.Insert(10, 10)
	BT.Insert(11, 11)
	BT.Insert(9, 9)
	BT.Insert(8, 8)
	BT.Insert(14, 14)
	BT.Insert(12, 12)
	BT.Insert(13, 13)
	t.Log(BT.String())

	expected := "[10 12 ] \n[8 9 ] [11 ] [13 14 ] \nX X X X X \n"
	actual := BT.String()
	if actual != expected {
		t.Fatal("Expected output: \n" + expected + "\n but got \n" + BT.String())
	}
}
