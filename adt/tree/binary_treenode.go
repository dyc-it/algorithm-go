package tree

import (
	"fmt"
	"algorithm-go/adt/queue"
	"algorithm-go/adt"
	"container/list"
	"algorithm-go/adt/stack"
)

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Value interface{}
}

/*
二叉树的递归先序遍历
 */
func (node *BinaryTreeNode) PreOrderRecursive() {
	fmt.Print(node.Value)
	if node.Left != nil {
		node.Left.PreOrderRecursive()
	}
	if node.Right != nil {
		node.Right.PreOrderRecursive()
	}
}


/*
二叉树的非递归先序遍历
思路:
1)初始情况先将节点N入栈
2)将节点N出栈放入队列Q中,然后依次将节点N的右节点RN和左节点LN入栈
3)出栈,重复2),栈为空时停止

TODO:写几个demo,便于理解
 */
func (node *BinaryTreeNode) PreOrder() (ret *list.List) {
	if node == nil {
		return nil
	}

	ret = list.New()
	s := stack.MakeStack()
	s.Push(&adt.Node{node})

	for !s.IsEmpty() {
		currentNode := s.Top().Value.(*BinaryTreeNode)
		//fmt.Println(tmpNode.Value)
		ret.PushBack(currentNode.Value)
		s.Pop()
		if currentNode.Right != nil {
			s.Push(&adt.Node{currentNode.Right})
		}
		if currentNode.Left != nil {
			s.Push(&adt.Node{currentNode.Left})
		}
	}
	return ret
}

/*
二叉树的递归中序遍历
 */
func (node *BinaryTreeNode) InOrderRecursive() {
	if node.Left != nil {
		node.Left.InOrderRecursive()
	}
	fmt.Print(node.Value)
	if node.Right != nil {
		node.Right.InOrderRecursive()
	}

}

/*
二叉树的非递归中序遍历
思路:
初始情况:设置根节点为当前节点N
当前节点N入栈
如果节点N的左子节点不为空,则将节点N入栈;并将节点N的左子节点设置为当前节点,对当前节点做同样的处理(入栈)
如果节点N的左子节点为空,取栈顶元素并将当前节点N设置为栈顶节点的右子节点,对栈顶元素进行出栈操作(出栈)

直到栈为空并且当前节点N为nil

TODO:写几个demo,便于理解
 */
func (node *BinaryTreeNode) InOrder() (ret *list.List) {
	if node == nil {
		return nil
	}

	ret = list.New()
	s := stack.MakeStack()

	currentNode := node

	for currentNode != nil || !s.IsEmpty() {
		for currentNode != nil {
			s.Push(&adt.Node{currentNode})
			currentNode = currentNode.Left
		}

		if !s.IsEmpty() {
			tmpNode := s.Top().Value.(*BinaryTreeNode)
			ret.PushBack(tmpNode.Value)
			currentNode = tmpNode.Right
			s.Pop()
		}
	}

	return ret
}

func (node *BinaryTreeNode) PostOrderRecursive() {
	if node.Left != nil {
		node.Left.PostOrderRecursive()
	}
	if node.Right != nil {
		node.Right.PostOrderRecursive()
	}
	fmt.Print(node.Value)
}

/*
二叉树的非递归后序遍历
使用两个栈:
将根节点压入第一个栈
从第一个栈中弹出一个元素，压入第二个栈
然后分别将该节点的左右孩子压入第一个栈
重复步骤2和步骤3直到第一个栈为空
执行结束，第二个栈中就保存了所有节点的后序遍历输出结果。依次将元素从第二个栈中弹出即可。


TODO:写几个demo,便于理解
 */
func (node *BinaryTreeNode) PostOrder() (ret *list.List) {
	if node == nil {
		return nil
	}

	ret = list.New()
	s := stack.MakeStack()
	output := stack.MakeStack()

	s.Push(&adt.Node{node})
	for !s.IsEmpty() {
		currentNode := s.Top().Value.(*BinaryTreeNode)
		output.Push(&adt.Node{currentNode})

		s.Pop()
		if currentNode.Left != nil {
			s.Push(&adt.Node{currentNode.Left})
		}
		if currentNode.Right != nil {
			s.Push(&adt.Node{currentNode.Right})
		}

	}

	for !output.IsEmpty() {
		ret.PushBack(output.Pop().Value.(*BinaryTreeNode).Value)
	}

	return ret
}

func (node *BinaryTreeNode) LevelOrder() (ret *list.List) {
	ret = list.New()
	q := queue.MakeQueue()
	for node != nil {
		ret.PushBack(node.Value)

		if node.Left != nil {
			q.Add(&adt.Node{node.Left})
		}
		if node.Right != nil {
			q.Add(&adt.Node{node.Right})
		}

		if q.IsEmpty() {
			node = nil
		} else {
			node = q.Delete().Value.(*BinaryTreeNode)
		}
	}
	return ret
}