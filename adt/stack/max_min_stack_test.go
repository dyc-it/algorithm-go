package stack

import (
	"testing"
)

var node1_max_min Node = Node{value:1}
var node2_max_min Node = Node{value:2}
var node3_max_min Node = Node{value:3}
var node4_max_min Node = Node{value:4}
var node5_max_min Node = Node{value:5}

func TestMaxMinStack_Push(t *testing.T) {
	s := MakeMaxMinStack()

	s.Push(&node1_max_min)
	s.Push(&node2_max_min)
	s.Push(&node3_max_min)
	s.Push(&node4_max_min)
	s.Push(&node5_max_min)

	if s.Top().value != node5_max_min.value {
		t.Errorf("Testing fails")
	}
}

func TestMaxMinStack_Pop(t *testing.T) {
	s := MakeMaxMinStack()

	s.Push(&node1_max_min)
	s.Push(&node2_max_min)
	s.Push(&node3_max_min)

	s.Pop()
	s.Pop()

	if s.Top().value != node1_max_min.value {
		t.Errorf("Testing fails")
	}
}

func TestMaxMinStack_Min(t *testing.T) {
	s := MakeMaxMinStack()

	s.Push(&node3_max_min)
	s.Push(&node5_max_min)
	s.Push(&node1_max_min)
	s.Pop()
	t.Log(s.Min().value)
	if s.Min().value != node3_max_min.value {
		t.Errorf("Testing fails")
	}

}

func TestMaxMinStack_Max(t *testing.T) {
	s := MakeMaxMinStack()

	s.Push(&node3_max_min)
	s.Push(&node1_max_min)
	s.Push(&node5_max_min)
	s.Pop()

	if s.Max().value != node3_max_min.value {
		t.Errorf("Testing fails")
	}
}











