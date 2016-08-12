package tree

import (
	"container/list"
	"fmt"
	"testing"
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

/*
判断是遍历的结果是否和期望值一直,不一致则输出异常信息
*/
func assert(expected []string, ret *list.List, t *testing.T) {
	i := 0
	length := len(expected)
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

func TestPreOrderRecursive(t *testing.T) {
	nodeA.PreOrderRecursive()
	fmt.Println()
}

func TestPreOrderNonRecursive(t *testing.T) {
	expected := []string{"A", "B", "D", "E", "G", "H", "C", "F", "I"}
	ret := nodeA.PreOrder()
	assert(expected, ret, t)
}

func TestInOrder(t *testing.T) {
	nodeA.InOrderRecursive()
	fmt.Println()
}

func TestInOrderNonRecursive(t *testing.T) {
	expected := []string{"D", "B", "G", "E", "H", "A", "C", "I", "F"}
	ret := nodeA.InOrder()
	assert(expected, ret, t)
}

func TestPostOrder(t *testing.T) {
	nodeA.PostOrderRecursive()
	fmt.Println()
}

func TestPostOrderNonRecursive(t *testing.T) {
	expected := []string{"D", "G", "H", "E", "B", "I", "F", "C", "A"}
	ret := nodeA.PostOrder()
	assert(expected, ret, t)
}

func TestLevelOrder(t *testing.T) {
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	ret := nodeA.LevelOrder()
	assert(expected, ret, t)
}
