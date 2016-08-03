package stack

import "algorithm-go/adt"

type Stack struct {
	nodes []*adt.Node
	size  int
}

func MakeStack() *Stack {
	return &Stack{nodes:make([]*adt.Node, 0)}
}

func (s *Stack) Push(node *adt.Node) {
	s.nodes = append(s.nodes, node)
	s.size++
}

func (s *Stack) Pop() *adt.Node {
	if s.size == 0 {
		return nil
	}
	s.size--
	return s.nodes[s.size - 1]
}

func (s *Stack) Top() *adt.Node {
	if s.size == 0 {
		return nil
	}
	return s.nodes[s.size - 1]
}


