package tree

import (
	"testing"
)

// The structure of example tree
//               A
//          B         C
//       D     E          F
//            G H        I

var binaryTreeNodeG BinaryTreeNode = BinaryTreeNode{nil, nil, "G"}
var binaryTreeNodeH BinaryTreeNode = BinaryTreeNode{nil, nil, "H"}
var binaryTreeNodeI BinaryTreeNode = BinaryTreeNode{nil, nil, "I"}

var binaryTreeNodeF BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeI, nil, "F"}

var binaryTreeNodeD BinaryTreeNode = BinaryTreeNode{nil, nil, "D"}
var binaryTreeNodeE BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeG, &binaryTreeNodeH, "E"}

var binaryTreeNodeB BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeD, &binaryTreeNodeE, "B"}

var binaryTreeNodeC BinaryTreeNode = BinaryTreeNode{nil, &binaryTreeNodeF, "C"}

var binaryTreeNodeA BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeB, &binaryTreeNodeC, "A"}

var binaryTree BinaryTree = BinaryTree{&binaryTreeNodeA}

func TestBinaryTreePreOrder(t *testing.T) {
	expected := []string{"A", "B", "D", "E", "G", "H", "C", "F", "I"}
	ret := binaryTree.PreOrder()
	assert(expected, ret, t)
}

func TestBinaryTreeInOrder(t *testing.T) {
	expected := []string{"D", "B", "G", "E", "H", "A", "C", "I", "F"}
	ret := binaryTree.InOrder()
	assert(expected, ret, t)
}

func TestBinaryTreePostOrder(t *testing.T) {
	expected := []string{"D", "G", "H", "E", "B", "I", "F", "C", "A"}
	ret := binaryTree.PostOrder()
	assert(expected, ret, t)
}

func TestBinaryTreeLevelOrder(t *testing.T) {
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	ret := binaryTree.LevelOrder()
	assert(expected, ret, t)
}
