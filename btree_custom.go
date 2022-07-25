package btree

import (
	"fmt"
	"os"
)

// NearestFunc[T] determines if newKey is nearer than tempNearest
// if it is nearer, return true
type NearestFunc[T any] func(goal, newKey, tempNearest T) bool

// get finds the given key in the subtree and returns it.
func (n *node[T]) getByFuncNearest(key T, nearest T, nearestFunc NearestFunc[T]) (_ T, _ bool) {

	var nearestInItems int
	for i, item := range n.items {
		if !n.cow.less(item, key) {
			nearestInItems = i
			break
		}
	}

	if !n.cow.less(key, n.items[nearestInItems]) {
		return n.items[nearestInItems], true
	}

	fmt.Printf("nearest one is [%v] %v %v %v\n", nearestInItems, n.items[nearestInItems])

	// i, found := n.items.find(key, n.cow.less)
	// if found {
	// 	nearest = n.items[i]
	// 	return n.items[i], true
	// } else

	if len(n.children) > 0 {

		fmt.Printf("%v %v %v %v\n", n.items[nearestInItems], nearestInItems, len(n.children), len(n.items))
		if nearestFunc(key, n.items[nearestInItems], nearest) {
			return n.children[nearestInItems+1].getByFuncNearest(key, n.items[nearestInItems], nearestFunc)
		}

		// if temp node is not nearer, it's child nodes wouldn't be nearer anymore
		return n.children[nearestInItems].getByFuncNearest(key, nearest, nearestFunc)
	}

	return nearest, false
}

// Get looks for the key item in the tree, returning it.  It returns
// (zeroValue, false) if unable to find that item.
func (t *BTreeG[T]) GetByFuncNearest(key T, nearest T, nearestFunc NearestFunc[T]) (_ T, _ bool) {
	if t.root == nil {
		return
	}
	t.root.print(os.Stdout, 0)
	return t.root.getByFuncNearest(key, nearest, nearestFunc)
}
