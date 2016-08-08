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

func (tree *Tree) Pre_order() {
	fmt.Print(tree.Value)
	if tree.Left != nil {
		tree.Left.Pre_order()
	}
	if tree.Right != nil {
		tree.Right.Pre_order()
	}
}

func (tree *Tree) In_order() {
	if tree.Left != nil {
		tree.Left.In_order()
	}
	fmt.Print(tree.Value)
	if tree.Right != nil {
		tree.Right.In_order()
	}

}

func (tree *Tree) Post_order() {
	if tree.Left != nil {
		tree.Left.Post_order()
	}
	if tree.Right != nil {
		tree.Right.Post_order()
	}
	fmt.Print(tree.Value)
}

func (tree *Tree) Level_order() {
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







