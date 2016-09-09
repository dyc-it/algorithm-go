package tree

import (
	"testing"
	"container/list"
)

// The structure of example tree


var arrayForMakingTree = []int{1, 0, 2, 3, 9, 4, 5, 8, 6, 7}
var tree BinarySearchTree = BinarySearchTree{nil}

func assertInt(expected []int, ret *list.List, t *testing.T) {
	i := 0
	length := len(expected)
	// 如果返回的list长度为0,需要加入判断;否则后面的for语句无法执行
	if ret.Len() < length {
		t.Errorf("tranverse failed!\t The amount of expected nodes is more than tranverse")
	}

	for iter := ret.Front(); iter != nil; iter = iter.Next() {
		//fmt.Print(iter.Value)
		if i >= length {
			t.Errorf("tranverse failed!\t The amount of expected nodes is less than tranverse")
		}
		if i < length && expected[i] != iter.Value {
			t.Errorf("Level order tranverse failed!\texpected: %s, get: %s", expected[i], iter.Value)
		}
		i++
	}
}

func TestBinarySearchTreeInsert(t *testing.T) {
	tree.makeBSTreeFromArray(arrayForMakingTree)

	ret := tree.InOrder()
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assertInt(expected, ret, t)
}

func TestBinarySearchTreePreOrder(t *testing.T) {
	tree.makeBSTreeFromArray(arrayForMakingTree)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ret := tree.PreOrder()
	assertInt(expected, ret, t)
}

func TestBinarySearchTreeInOrder(t *testing.T) {
	tree.makeBSTreeFromArray(arrayForMakingTree)
	expected := []int{1}
	ret := tree.InOrder()
	assertInt(expected, ret, t)
}

func TestBinarySearchTreePostOrder(t *testing.T) {
	tree.makeBSTreeFromArray(arrayForMakingTree)
	expected := []int{4, 7, 8, 5, 2, 9, 6, 3, 1}
	ret := tree.PostOrder()
	assertInt(expected, ret, t)
}

func TestBinarySearchTreeLevelOrder(t *testing.T) {
	tree.makeBSTreeFromArray(arrayForMakingTree)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ret := tree.LevelOrder()
	assertInt(expected, ret, t)
}
