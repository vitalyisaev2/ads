// Skiena 3.4: Binary Search Trees
// Simple thread-safe binary search tree

package structures

import (
	"fmt"
	"strings"
	"sync"
)

// BinarySearchTreeItem describes interfaces for the BinarySearchTree payload
type BinarySearchTreeItem interface {
	Less(interface{}) bool
	Equal(interface{}) bool
	//Value() interface{}
}

// BinarySearchTree is a simple implementation of a binary search tree
type BinarySearchTree interface {
	Search(BinarySearchTreeItem) BinarySearchTreeItem
	Insert(BinarySearchTreeItem)
	Min() BinarySearchTreeItem
	Max() BinarySearchTreeItem
	Root() BinarySearchTreeItem
	Len() int
	Items() []BinarySearchTreeNode
}

type binarySearchTreeImpl struct {
	sync.RWMutex
	root *binarySearchTreeNodeImpl
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

func (tree *binarySearchTreeImpl) Max() BinarySearchTreeItem {
	tree.RLock()
	defer tree.RUnlock()
	return tree.root.max()
}

func (tree *binarySearchTreeImpl) Len() int {
	tree.RLock()
	defer tree.RUnlock()
	return tree.size
}

func (tree *binarySearchTreeImpl) Items() []BinarySearchTreeNode {
	tree.RLock()
	defer tree.RUnlock()
	items := make([]BinarySearchTreeNode, 0, tree.size)
	tree.root.traverseSlice(&items)
	return items
}

func (tree *binarySearchTreeImpl) Root() BinarySearchTreeItem {
	tree.RLock()
	defer tree.RUnlock()
	return tree.root.item
}

func (tree *binarySearchTreeImpl) String() string {
	tree.RLock()
	defer tree.RUnlock()
	items := tree.Items()
	itemsString := make([]string, 0, cap(items))
	for i := range items {
		itemsString = append(itemsString, items[i].String())
	}
	return strings.Join(itemsString, "\n")
}

// BinarySearchTreeNode contains a payload and also holds links to the
// parents and children
type BinarySearchTreeNode interface {
	fmt.Stringer
	Item() BinarySearchTreeItem
	Parent() BinarySearchTreeNode
	Left() BinarySearchTreeNode
	Right() BinarySearchTreeNode
}

type binarySearchTreeNodeImpl struct {
	item   BinarySearchTreeItem
	parent *binarySearchTreeNodeImpl
	left   *binarySearchTreeNodeImpl
	right  *binarySearchTreeNodeImpl
}

func (node *binarySearchTreeNodeImpl) Item() BinarySearchTreeItem   { return node.item }
func (node *binarySearchTreeNodeImpl) Parent() BinarySearchTreeNode { return node.parent }
func (node *binarySearchTreeNodeImpl) Left() BinarySearchTreeNode   { return node.left }
func (node *binarySearchTreeNodeImpl) Right() BinarySearchTreeNode  { return node.right }

func (node *binarySearchTreeNodeImpl) insert(newItem BinarySearchTreeItem) {
	if node.item == nil {
		node.item = newItem
		return
	}

	if node.item.Equal(newItem) {
		return
	}

	if node.item.Less(newItem) {
		if node.right == nil {
			newNode := &binarySearchTreeNodeImpl{
				item:   newItem,
				parent: node,
			}
			node.right = newNode
			return
		}
		node.right.insert(newItem)
		return
	}

	if node.left == nil {
		newNode := &binarySearchTreeNodeImpl{
			item:   newItem,
			parent: node,
		}
		node.left = newNode
		return
	}
	node.left.insert(newItem)
}

func (node *binarySearchTreeNodeImpl) search(soughtItem BinarySearchTreeItem) BinarySearchTreeItem {
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

func (node *binarySearchTreeNodeImpl) min() BinarySearchTreeItem {
	minNode := node
	for minNode.left != nil {
		minNode = minNode.left
	}
	return minNode.item
}

func (node *binarySearchTreeNodeImpl) max() BinarySearchTreeItem {
	maxNode := node
	for maxNode.right != nil {
		maxNode = maxNode.right
	}
	return maxNode.item
}

func (node *binarySearchTreeNodeImpl) traverseSlice(items *[]BinarySearchTreeNode) {
	if node.item != nil {
		if node.left != nil {
			node.left.traverseSlice(items)
		}
		*items = append(*items, node)
		if node.right != nil {
			node.right.traverseSlice(items)
		}
	}
}

func (node *binarySearchTreeNodeImpl) String() string {
	var (
		item   string
		parent string
		left   string
		right  string
	)
	item = fmt.Sprintf("%v", node.item)
	switch node.parent {
	case nil:
		parent = "nil"
	default:
		parent = fmt.Sprintf("%v", node.parent.item)
	}
	switch node.left {
	case nil:
		left = "nil"
	default:
		left = fmt.Sprintf("%v", node.left.item)
	}
	switch node.right {
	case nil:
		right = "nil"
	default:
		right = fmt.Sprintf("%v", node.right.item)
	}
	return fmt.Sprintf(
		"item=%s(%p) parent=%s left=%s right=%s",
		item, node, parent, left, right)
}

// NewBinarySearchTree returns empty binary search tree
func NewBinarySearchTree() BinarySearchTree {
	root := &binarySearchTreeNodeImpl{}
	tree := &binarySearchTreeImpl{root: root, size: 0}
	return tree
}
