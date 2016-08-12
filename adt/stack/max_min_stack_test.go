package stack

import (
	"algorithm-go/adt"
	"testing"
)

var node1_max_min adt.Node = adt.Node{Value: 1}
var node2_max_min adt.Node = adt.Node{Value: 2}
var node3_max_min adt.Node = adt.Node{Value: 3}
var node4_max_min adt.Node = adt.Node{Value: 4}
var node5_max_min adt.Node = adt.Node{Value: 5}

func TestMaxMinStack_Push(t *testing.T) {
	s := MakeMaxMinStack()

	s.Push(&node1_max_min)
	s.Push(&node2_max_min)
	s.Push(&node3_max_min)
	s.Push(&node4_max_min)
	s.Push(&node5_max_min)

	if s.Top().Value != node5_max_min.Value {
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

	if s.Top().Value != node1_max_min.Value {
		t.Errorf("Testing fails")
	}
}

func TestMaxMinStack_Min(t *testing.T) {
	s := MakeMaxMinStack()

	s.Push(&node3_max_min)
	s.Push(&node5_max_min)
	s.Push(&node1_max_min)
	s.Pop()
	t.Log(s.Min().Value)
	if s.Min().Value != node3_max_min.Value {
		t.Errorf("Testing fails")
	}

}

func TestMaxMinStack_Max(t *testing.T) {
	s := MakeMaxMinStack()

	s.Push(&node3_max_min)
	s.Push(&node1_max_min)
	s.Push(&node5_max_min)
	s.Pop()

	if s.Max().Value != node3_max_min.Value {
		t.Errorf("Testing fails")
	}
}
