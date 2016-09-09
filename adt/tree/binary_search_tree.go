package tree

import "container/list"

type BinarySearchTree struct {
	root *BinaryTreeNode
}

func (tree *BinarySearchTree) PreOrder() (ret *list.List) {
	return tree.root.PreOrder()
}

func (tree *BinarySearchTree) InOrder() (ret *list.List) {
	return tree.root.InOrder()
}

func (tree *BinarySearchTree) PostOrder() (ret *list.List) {
	return tree.root.PostOrder()
}

func (tree *BinarySearchTree) LevelOrder() (ret *list.List) {
	return tree.root.LevelOrder()
}

// 二叉查找树的方法:插入、删除、搜索

/*
二叉查找树的插入:
1.找到待插入节点的父节点位置
2.将待插入节点的父节点设置为插入的位置
3.判断带插入节点是其父节点的左子节点还是右子节点
*/
func (tree *BinarySearchTree) insert(insertedNode *BinaryTreeNode) {

	var parentNode *BinaryTreeNode = nil

	var tmpNode *BinaryTreeNode = tree.root
	for tmpNode != nil {
		parentNode = tmpNode
		if insertedNode.Value.(int) < tmpNode.Value.(int) {
			tmpNode = tmpNode.Left
		} else {
			tmpNode = tmpNode.Right
		}
	}

	insertedNode.Parent = parentNode

	if parentNode == nil {
		tree.root = insertedNode
	} else if insertedNode.Value.(int) < parentNode.Value.(int) {
		parentNode.Left = insertedNode
	} else {
		parentNode.Right = insertedNode
	}
}

func (tree *BinarySearchTree) makeBSTreeFromArray(array []int) *BinarySearchTree {
	for _, value := range array {
		var node BinaryTreeNode = BinaryTreeNode{Left: nil, Right: nil, Parent: nil, Value: value}
		tree.insert(&node)
	}
	return tree
}
