package btree

import "fmt"

// NearestFunc[T] determines if newKey is nearer than tempNearest
// if it is nearer, return true
type NearestFunc[T any] func(goal, newKey, tempNearest T) bool

// get finds the given key in the subtree and returns it.
func (n *node[T]) getByFuncNearest(key T, nearest T, nearestFunc NearestFunc[T]) (_ T, _ bool) {
	i, found := n.items.find(key, n.cow.less)
	if found {
		nearest = n.items[i]
		return n.items[i], true
	} else if len(n.children) > 0 && len(n.children) > i {

		fmt.Printf("%v %v %v", n.items[i], i, len(n.children))
		if nearestFunc(key, n.items[i], nearest) {
			return n.children[i].getByFuncNearest(key, n.items[i], nearestFunc)
		}

		// if temp node is not nearer, it's child nodes wouldn't be nearer anymore
		return n.children[i].getByFuncNearest(key, nearest, nearestFunc)
	}

	return nearest, false
}

// Get looks for the key item in the tree, returning it.  It returns
// (zeroValue, false) if unable to find that item.
func (t *BTreeG[T]) GetByFuncNearest(key T, nearest T, nearestFunc NearestFunc[T]) (_ T, _ bool) {
	if t.root == nil {
		return
	}
	return t.root.getByFuncNearest(key, nearest, nearestFunc)
}
