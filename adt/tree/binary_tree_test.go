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

func TestBinaryTreeNodePre_order(t *testing.T) {
	binaryTree.root.PreOrderRecursive()
	fmt.Println()
}

func TestBinaryTreeNodeIn_order(t *testing.T) {
	binaryTreeNodeA.InOrderRecursive()
	fmt.Println()
}

func TestBinaryTreeNodePost_order(t *testing.T) {
	binaryTreeNodeA.PostOrderRecursive()
	fmt.Println()
}

func TestBinaryTreeNodeLevel_order(t *testing.T) {
	binaryTreeNodeA.LevelOrder()
	fmt.Println()
}


