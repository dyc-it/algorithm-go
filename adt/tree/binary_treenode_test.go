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

var nodeG BinaryTreeNode = BinaryTreeNode{nil, nil, "G"}
var nodeH BinaryTreeNode = BinaryTreeNode{nil, nil, "H"}
var nodeI BinaryTreeNode = BinaryTreeNode{nil, nil, "I"}

var nodeF BinaryTreeNode = BinaryTreeNode{&nodeI, nil, "F"}

var nodeD BinaryTreeNode = BinaryTreeNode{nil, nil, "D"}
var nodeE BinaryTreeNode = BinaryTreeNode{&nodeG, &nodeH, "E"}

var nodeB BinaryTreeNode = BinaryTreeNode{&nodeD, &nodeE, "B"}

var nodeC BinaryTreeNode = BinaryTreeNode{nil, &nodeF, "C"}

var nodeA BinaryTreeNode = BinaryTreeNode{&nodeB, &nodeC, "A"}

func TestPreOrderRecursive(t *testing.T) {
	nodeA.PreOrderRecursive()
	fmt.Println()
}

func TestPreOrderNonRecursive(t *testing.T) {
	i := 0
	expected := []string{"A", "B", "D", "E", "G", "H", "C", "F", "I"}
	length := len(expected)

	ret := nodeA.PreOrder()
	for iter := ret.Front(); iter != nil; iter = iter.Next() {
		//fmt.Print(iter.Value)
		if i >= length {
			t.Errorf("Level order tranverse failed!\t The amount of expected nodes is less than tranverse")
		}
		if i < length && expected[i] != iter.Value {
			t.Errorf("Level order tranverse failed!\texpected: %s, get: %s", expected[i], iter.Value)
		}
		i++
	}
}

func TestInOrder(t *testing.T) {
	nodeA.InOrderRecursive()
	fmt.Println()
}

func TestInOrderNonRecursive(t *testing.T) {
	i := 0
	expected := []string{"D", "B", "G", "E", "H", "A", "C", "I", "F"}
	length := len(expected)

	ret := nodeA.InOrder()

	// 如果返回的list长度为0,需要加入判断;否则后面的for语句无法执行
	if ret.Len() < length {
		t.Errorf("Level order tranverse failed!\t The amount of expected nodes is more than tranverse")
	}

	for iter := ret.Front(); iter != nil; iter = iter.Next() {
		//fmt.Print(iter.Value)
		if i >= length {
			t.Errorf("Level order tranverse failed!\t The amount of expected nodes is less than tranverse")
		}
		if i < length && expected[i] != iter.Value {
			t.Errorf("Level order tranverse failed!\texpected: %s, get: %s", expected[i], iter.Value)
		}
		i++
	}
}

func TestPostOrder(t *testing.T) {
	nodeA.PostOrderRecursive()
	fmt.Println()
}

func TestPostOrderNonRecursive(t *testing.T) {
	i := 0
	expected := []string{"D", "G", "H", "E", "B", "I", "F", "C", "A"}
	length := len(expected)

	ret := nodeA.PostOrder()

	// 如果返回的list长度为0,需要加入判断;否则后面的for语句无法执行
	if ret.Len() < length {
		t.Errorf("Level order tranverse failed!\t The amount of expected nodes is more than tranverse")
	}

	for iter := ret.Front(); iter != nil; iter = iter.Next() {
		//fmt.Print(iter.Value)
		if i >= length {
			t.Errorf("Level order tranverse failed!\t The amount of expected nodes is less than tranverse")
		}
		if i < length && expected[i] != iter.Value {
			t.Errorf("Level order tranverse failed!\texpected: %s, get: %s", expected[i], iter.Value)
		}
		i++
	}
}

func TestLevelOrder(t *testing.T) {
	i := 0
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	length := len(expected)

	ret := nodeA.LevelOrder()

	// 如果返回的list长度为0,需要加入判断;否则后面的for语句无法执行
	if ret.Len() < length {
		t.Errorf("Level order tranverse failed!\t The amount of expected nodes is more than tranverse")
	}

	for iter := ret.Front(); iter != nil; iter = iter.Next() {
		//fmt.Print(iter.Value)
		if i >= length {
			t.Errorf("Level order tranverse failed!\t The amount of expected nodes is less than tranverse")
		}
		if i < length && expected[i] != iter.Value {
			t.Errorf("Level order tranverse failed!\texpected: %s, get: %s", expected[i], iter.Value)
		}
		i++
	}
}


