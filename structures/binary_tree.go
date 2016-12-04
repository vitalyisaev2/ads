// Skiena 3.4: Binary Search Trees
// Simple thread-safe binary search tree

package structures

import (
	"sync"
)

// BinarySearchTreeItem describes interfaces for the BinarySearchTree payload
type BinarySearchTreeItem interface {
	Less(BinarySearchTreeItem) bool
	Equal(BinarySearchTreeItem) bool
}

// BinarySearchTree is a simple implementation of a binary search tree
type BinarySearchTree interface {
	Search(BinarySearchTreeItem) BinarySearchTreeNode
	//Insert(BinarySearchTreeItem)
	Min() BinarySearchTreeItem
	//Max() BinarySearchTreeItem
	//Traverse() []BinarySearchTreeItem
	Len() int
}

type binarySearchTreeImpl struct {
	sync.RWMutex
	root *binarySearchTreeNodeImpl
	size int
}

func (tree *binarySearchTreeImpl) Search(sought BinarySearchTreeItem) BinarySearchTreeNode {
	tree.RLock()
	defer tree.RUnlock()
	return tree.root.search(sought)
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

// BinarySearchTreeNode represents node of a binary search tree
type BinarySearchTreeNode interface {
	Item() BinarySearchTreeItem
	search(BinarySearchTreeItem) BinarySearchTreeNode
	min() BinarySearchTreeItem
}

type binarySearchTreeNodeImpl struct {
	item   BinarySearchTreeItem
	parent *binarySearchTreeNodeImpl
	left   *binarySearchTreeNodeImpl
	right  *binarySearchTreeNodeImpl
}

func (node *binarySearchTreeNodeImpl) Item() BinarySearchTreeItem { return node.item }

func (node *binarySearchTreeNodeImpl) search(sought BinarySearchTreeItem) BinarySearchTreeNode {
	if node.item == nil {
		return nil
	}

	if node.item.Equal(sought) {
		return node
	}

	if node.item.Less(sought) {
		if node.right != nil {
			return node.right.search(sought)
		}
	} else {
		if node.left != nil {
			return node.left.search(sought)
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

// NewBinarySearchTree returns empty binary search tree
func NewBinarySearchTree() BinarySearchTree {
	root := &binarySearchTreeNodeImpl{}
	tree := &binarySearchTreeImpl{root: root, size: 0}
	return tree
}
