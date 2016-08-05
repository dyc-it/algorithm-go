package tree

import (
	"fmt"
	"algorithm-go/adt/queue"
	"algorithm-go/adt"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Value interface{}
}

func (tree *Tree) pre_order() {
	fmt.Print(tree.Value)
	if tree.Left != nil {
		tree.Left.pre_order()
	}
	if tree.Right != nil {
		tree.Right.pre_order()
	}
}

func (tree *Tree) in_order() {
	if tree.Left != nil {
		tree.Left.in_order()
	}
	fmt.Print(tree.Value)
	if tree.Right != nil {
		tree.Right.in_order()
	}

}

func (tree *Tree) post_order() {
	if tree.Left != nil {
		tree.Left.post_order()
	}
	if tree.Right != nil {
		tree.Right.post_order()
	}
	fmt.Print(tree.Value)
}

func (tree *Tree) level_order() {
	q := queue.MakeQueue()
	for tree != nil {
		fmt.Print(tree.Value)

		if tree.Left != nil {
			q.Add(&adt.Node{tree.Left})
		}
		if tree.Right != nil {
			q.Add(&adt.Node{tree.Right})
		}

		if q.Is_empty() {
			tree = nil
		} else {
			tree = q.Delete().Value.(*Tree)
		}
	}

}







