package tree

import (
	"testing"
	"fmt"
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

func TestBinarySearchTreePre_order(t *testing.T) {
	binarySearchTree.root.PreOrderRecursive()
	fmt.Println()
}

func TestBinarySearchTreeIn_order(t *testing.T) {
	binarySearchTreeNodeA.InOrderRecursive()
	fmt.Println()
}

func TestBinarySearchTreePost_order(t *testing.T) {
	binarySearchTreeNodeA.PostOrderRecursive()
	fmt.Println()
}

func TestBinarySearchTreeNodeLevel_order(t *testing.T) {
	binarySearchTreeNodeA.LevelOrder()
	fmt.Println()
}


