package stack

import "testing"


func TestStack_Push(t *testing.T) {
	var node1 Node = Node{value:1}
	var node2 Node = Node{value:2}
	var node3 Node = Node{value:3}
	var node4 Node = Node{value:4}
	var node5 Node = Node{value:5}

	s := MakeStack()

	s.Push(&node1)
	s.Push(&node2)
	s.Push(&node3)
	s.Push(&node4)
	s.Push(&node5)

	if s.Top().value != node5.value {
		t.Errorf("Testing fails")
	}
}

func TestStack_Pop(t *testing.T) {
	var node1 Node = Node{value:1}
	var node2 Node = Node{value:2}
	var node3 Node = Node{value:3}

	s := MakeStack()
	s.Push(&node1)
	s.Push(&node2)
	s.Push(&node3)

	s.Pop()
	s.Pop()

	if s.Top().value != node1.value {
		t.Errorf("Testing fails")
	}
}
