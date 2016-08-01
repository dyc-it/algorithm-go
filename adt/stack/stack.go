package stack


type Stack struct {
	nodes []*Node
	size int
}

func MakeStack() *Stack {
	return &Stack {nodes:make([]*Node, 0)}
}

func (s *Stack) Push(node *Node) {
	s.nodes = append(s.nodes, node)
	s.size++
}

func (s *Stack) Pop() *Node {
	if s.size == 0 {
		return nil
	}
	s.size--
	return s.nodes[s.size-1]
}

func (s *Stack) Top() *Node {
	if s.size == 0 {
		return nil
	}
	return s.nodes[s.size-1]
}


