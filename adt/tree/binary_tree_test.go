package tree

import (
	"testing"
)

// The structure of example tree
//               A
//          B         C
//       D     E          F
//            G H        I

var binaryTreeNodeG BinaryTreeNode = BinaryTreeNode{nil, nil, "G", nil}
var binaryTreeNodeH BinaryTreeNode = BinaryTreeNode{nil, nil, "H", nil}
var binaryTreeNodeI BinaryTreeNode = BinaryTreeNode{nil, nil, "I", nil}

var binaryTreeNodeF BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeI, nil, "F", nil}

var binaryTreeNodeD BinaryTreeNode = BinaryTreeNode{nil, nil, "D", nil}
var binaryTreeNodeE BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeG, &binaryTreeNodeH, "E", nil}

var binaryTreeNodeB BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeD, &binaryTreeNodeE, "B", nil}

var binaryTreeNodeC BinaryTreeNode = BinaryTreeNode{nil, &binaryTreeNodeF, "C", nil}

var binaryTreeNodeA BinaryTreeNode = BinaryTreeNode{&binaryTreeNodeB, &binaryTreeNodeC, "A", nil}

var binaryTree BinaryTree = BinaryTree{&binaryTreeNodeA}

func TestBinaryTreePreOrder(t *testing.T) {
	expected := []string{"A", "B", "D", "E", "G", "H", "C", "F", "I"}
	ret := binaryTree.PreOrder()
	assertString(expected, ret, t)
}

func TestBinaryTreeInOrder(t *testing.T) {
	expected := []string{"D", "B", "G", "E", "H", "A", "C", "I", "F"}
	ret := binaryTree.InOrder()
	assertString(expected, ret, t)
}

func TestBinaryTreePostOrder(t *testing.T) {
	expected := []string{"D", "G", "H", "E", "B", "I", "F", "C", "A"}
	ret := binaryTree.PostOrder()
	assertString(expected, ret, t)
}

func TestBinaryTreeLevelOrder(t *testing.T) {
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	ret := binaryTree.LevelOrder()
	assertString(expected, ret, t)
}
