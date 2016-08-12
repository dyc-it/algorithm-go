package tree

import (
	"testing"
)

// The structure of example tree
//               A
//          B         C
//       D     E          F
//            G H        I

var binarySearchTreeNodeG BinaryTreeNode = BinaryTreeNode{nil, nil, "G"}
var binarySearchTreeNodeH BinaryTreeNode = BinaryTreeNode{nil, nil, "H"}
var binarySearchTreeNodeI BinaryTreeNode = BinaryTreeNode{nil, nil, "I"}

var binarySearchTreeNodeF BinaryTreeNode = BinaryTreeNode{&binarySearchTreeNodeI, nil, "F"}

var binarySearchTreeNodeD BinaryTreeNode = BinaryTreeNode{nil, nil, "D"}
var binarySearchTreeNodeE BinaryTreeNode = BinaryTreeNode{&binarySearchTreeNodeG, &binarySearchTreeNodeH, "E"}

var binarySearchTreeNodeB BinaryTreeNode = BinaryTreeNode{&binarySearchTreeNodeD, &binarySearchTreeNodeE, "B"}

var binarySearchTreeNodeC BinaryTreeNode = BinaryTreeNode{nil, &binarySearchTreeNodeF, "C"}

var binarySearchTreeNodeA BinaryTreeNode = BinaryTreeNode{&binarySearchTreeNodeB, &binarySearchTreeNodeC, "A"}

var binarySearchTree BinaryTree = BinaryTree{&binarySearchTreeNodeA}

func TestBinarySearchTreePreOrder(t *testing.T) {
	expected := []string{"A", "B", "D", "E", "G", "H", "C", "F", "I"}
	ret := binarySearchTree.PreOrder()
	assert(expected, ret, t)
}

func TestBinarySearchTreeInOrder(t *testing.T) {
	expected := []string{"D", "B", "G", "E", "H", "A", "C", "I", "F"}
	ret := binarySearchTree.InOrder()
	assert(expected, ret, t)
}

func TestBinarySearchTreePostOrder(t *testing.T) {
	expected := []string{"D", "G", "H", "E", "B", "I", "F", "C", "A"}
	ret := binarySearchTree.PostOrder()
	assert(expected, ret, t)
}

func TestBinarySearchTreeLevelOrder(t *testing.T) {
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	ret := binarySearchTree.LevelOrder()
	assert(expected, ret, t)
}
