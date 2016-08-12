package stack

import (
	"algorithm-go/adt"
	"container/list"
)

type Stack struct {
	nodes []*adt.Node
	size  int
	list  *list.List
}

func MakeStack() *Stack {
	return &Stack{list: list.New()}
}

func (s *Stack) Push(node *adt.Node) {
	s.list.PushFront(node)
}

func (s *Stack) Pop() *adt.Node {
	if s.IsEmpty() {
		return nil
	} else {
		element := s.list.Front()
		s.list.Remove(element)
		node := element.Value.(*adt.Node)
		return node
	}
}

func (s *Stack) Top() *adt.Node {
	if s.IsEmpty() {
		return nil
	} else {
		node := s.list.Front().Value.(*adt.Node)
		return node
	}
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
