package stack

import (
	"testing"
	"algorithm-go/adt"
)

func TestStack_Push(t *testing.T) {
	var node1 adt.Node = adt.Node{Value:1}
	var node2 adt.Node = adt.Node{Value:2}
	var node3 adt.Node = adt.Node{Value:3}
	var node4 adt.Node = adt.Node{Value:4}
	var node5 adt.Node = adt.Node{Value:5}

	s := MakeStack()

	s.Push(&node1)
	s.Push(&node2)
	s.Push(&node3)
	s.Push(&node4)
	s.Push(&node5)

	if s.Top().Value != node5.Value {
		t.Errorf("Testing fails")
	}
}

func TestStack_Pop(t *testing.T) {
	var node1 adt.Node = adt.Node{Value:1}
	var node2 adt.Node = adt.Node{Value:2}
	var node3 adt.Node = adt.Node{Value:3}

	s := MakeStack()
	s.Push(&node1)
	s.Push(&node2)
	s.Push(&node3)

	s.Pop()
	s.Pop()

	if s.Top().Value != node1.Value {
		t.Errorf("Testing fails")
	}
}
