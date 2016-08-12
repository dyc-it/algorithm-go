package stack

import (
	"testing"
	"algorithm-go/adt"
)

var node1 adt.Node = adt.Node{Value:1}
var node2 adt.Node = adt.Node{Value:2}
var node3 adt.Node = adt.Node{Value:3}
var node4 adt.Node = adt.Node{Value:4}
var node5 adt.Node = adt.Node{Value:5}

func TestStack_Push(t *testing.T) {
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

func TestStack_IsEmpty(t *testing.T) {
	s := MakeStack()
	s.Push(&node1)
	s.Push(&node2)
	s.Push(&node3)

	s.Pop()
	s.Pop()

	// node1 remains in stack
	if s.IsEmpty() {
		t.Errorf("Testing fails")
	}

	// stack is empty
	s.Pop()
	if !s.IsEmpty() {
		t.Errorf("Testing fails")
	}

	// 是否处理了栈为空时,再弹出的情况
	s.Pop()
}

