package tree

import "container/list"

type BinaryTree struct {
	root *BinaryTreeNode
}

func (tree *BinaryTree) PreOrder() (ret *list.List) {
	return tree.root.PreOrder()
}

func (tree *BinaryTree) InOrder() (ret *list.List) {
	return tree.root.InOrder()
}

func (tree *BinaryTree) PostOrder() (ret *list.List) {
	return tree.root.PostOrder()
}

func (tree *BinaryTree) LevelOrder() (ret *list.List) {
	return tree.root.LevelOrder()
}
