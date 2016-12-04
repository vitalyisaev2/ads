// Skiena 3.4: Binary Search Trees
// Simple thread-safe binary search tree

package structures

import
//"fmt"

"sync"

// BinarySearchTreeItem describes interfaces for the BinarySearchTree payload
type BinarySearchTreeItem interface {
	Less(interface{}) bool
	Equal(interface{}) bool
}

// BinarySearchTree is a simple implementation of a binary search tree
type BinarySearchTree interface {
	//fmt.Stringer
	Search(BinarySearchTreeItem) BinarySearchTreeItem
	Insert(BinarySearchTreeItem)
	Min() BinarySearchTreeItem
	//Max() BinarySearchTreeItem
	Items() []BinarySearchTreeItem
	Root() BinarySearchTreeItem
	Len() int
}

type binarySearchTreeImpl struct {
	sync.RWMutex
	root *binarySearchTreeNode
	size int
}

func (tree *binarySearchTreeImpl) Insert(newItem BinarySearchTreeItem) {
	tree.Lock()
	defer tree.Unlock()
	tree.root.insert(newItem)
	tree.size++
}

func (tree *binarySearchTreeImpl) Search(soughtItem BinarySearchTreeItem) BinarySearchTreeItem {
	tree.RLock()
	defer tree.RUnlock()
	return tree.root.search(soughtItem)
}

func (tree *binarySearchTreeImpl) Min() BinarySearchTreeItem {
	tree.RLock()
	defer tree.RUnlock()
	return tree.root.min()
}

func (tree *binarySearchTreeImpl) Len() int {
	tree.RLock()
	defer tree.RUnlock()
	return tree.size
}

//func (tree *binarySearchTreeImpl) String() string {
//tree.RLock()
//defer tree.RUnlock()
//return fmt.Sprintf("Tree: total items %d", tree.size)
//}

func (tree *binarySearchTreeImpl) Items() []BinarySearchTreeItem {
	tree.RLock()
	defer tree.RUnlock()
	items := make([]BinarySearchTreeItem, 0, tree.size)
	tree.root.traverseSlice(&items)
	return items
}

func (tree *binarySearchTreeImpl) Root() BinarySearchTreeItem {
	tree.RLock()
	defer tree.RUnlock()
	return tree.root.item
}

type binarySearchTreeNode struct {
	item   BinarySearchTreeItem
	parent *binarySearchTreeNode
	left   *binarySearchTreeNode
	right  *binarySearchTreeNode
}

func (node *binarySearchTreeNode) Item() BinarySearchTreeItem { return node.item }

func (node *binarySearchTreeNode) insert(newItem BinarySearchTreeItem) {
	if node.item == nil {
		node.item = newItem
		return
	}

	if node.item.Equal(newItem) {
		return
	}

	if node.item.Less(newItem) {
		if node.right == nil {
			newNode := &binarySearchTreeNode{
				item:   newItem,
				parent: node,
			}
			node.right = newNode
			return
		}
		node.right.insert(newItem)
	}

	if node.left == nil {
		newNode := &binarySearchTreeNode{
			item:   newItem,
			parent: node,
		}
		node.left = newNode
		return
	}
	node.left.insert(newItem)
}

func (node *binarySearchTreeNode) search(soughtItem BinarySearchTreeItem) BinarySearchTreeItem {
	if node.item == nil {
		return nil
	}

	if node.item.Equal(soughtItem) {
		return node.item
	}

	if node.item.Less(soughtItem) {
		if node.right != nil {
			return node.right.search(soughtItem)
		}
	} else {
		if node.left != nil {
			return node.left.search(soughtItem)
		}
	}

	return nil
}

func (node *binarySearchTreeNode) min() BinarySearchTreeItem {
	minNode := node
	for minNode.left != nil {
		minNode = minNode.left
	}
	return minNode.item
}

func (node *binarySearchTreeNode) traverseSlice(items *[]BinarySearchTreeItem) {
	if node.item != nil {
		node.left.traverseSlice(items)
		*items = append(*items, node.item)
		node.right.traverseSlice(items)
	}
}

// NewBinarySearchTree returns empty binary search tree
func NewBinarySearchTree() BinarySearchTree {
	root := &binarySearchTreeNode{}
	tree := &binarySearchTreeImpl{root: root, size: 0}
	return tree
}
