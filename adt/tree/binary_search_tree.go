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
