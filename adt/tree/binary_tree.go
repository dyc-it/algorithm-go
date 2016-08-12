package tree

type BinaryTree struct {
	root *BinaryTreeNode
}

func (tree *BinaryTree) Pre_order() {
	tree.Pre_order()
}

func (tree *BinaryTree) In_order() {
	tree.In_order()
}

func (tree *BinaryTree) Post_order() {
	tree.Post_order()
}

func (tree *BinaryTree) Level_order() {
	tree.Level_order()
}






