package stack

import "algorithm-go/adt"

/*
	MaxMinStack provide max and min function to get the max and min value of the stack.
	To do so, three basic stack is needed; it cannot be done by variables, the reason if following:
	If we define a min variable to store the min value of stack,
	first pushing 3, 5, 1 into stack, and the pop 1 out of the stack
	stack		min
	3(push)		3
	35(push)	3
	351(push)	1
	35(pop)		?, obviously the value of top() is 5, and is not the min value

*/
type MaxMinStack struct {
	dataStack *Stack
	maxStack  *Stack
	minStack  *Stack
}

func MakeMaxMinStack() *MaxMinStack {
	return &MaxMinStack{dataStack: MakeStack(), maxStack: MakeStack(), minStack: MakeStack()}
}

func (s *MaxMinStack) Push(node *adt.Node) {
	s.dataStack.Push(node)

	// if maxStack is empty, means the dataStack has one node.
	if s.maxStack.IsEmpty() {
		s.maxStack.Push(node)
		s.minStack.Push(node)
	} else {
		// when minstack is not empty, if value is less than the top of minstack, push less value
		// to minstack, else push the top of minstack to minstack
		if s.minStack.Top().Value.(int) > node.Value.(int) {
			s.minStack.Push(node)
		} else {
			s.minStack.Push(s.minStack.Top())
		}

		// similar to minstack
		if s.maxStack.Top().Value.(int) < node.Value.(int) {
			s.maxStack.Push(node)
		} else {
			s.maxStack.Push(s.maxStack.Top())
		}
	}
}

func (s *MaxMinStack) Pop() *adt.Node {
	s.maxStack.Pop()
	s.minStack.Pop()
	return s.dataStack.Pop()
}

func (s *MaxMinStack) Top() *adt.Node {
	return s.dataStack.Top()
}

func (s *MaxMinStack) Max() *adt.Node {
	return s.maxStack.Top()
}

func (s *MaxMinStack) Min() *adt.Node {
	return s.minStack.Top()
}
