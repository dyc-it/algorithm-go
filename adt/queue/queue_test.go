package queue

import (
	"testing"
	"algorithm-go/adt"
)

func TestQueue_add(t *testing.T) {
	node1 := adt.Node{Value:1}
	node2 := adt.Node{Value:2}
	node3 := adt.Node{Value:3}

	q := MakeQueue()

	q.Add(&node1)
	q.Add(&node2)
	q.Add(&node3)
	if q.First().Value.(int) != node1.Value.(int) {
		t.Errorf("Testing fails: first value error")
	}

	if q.Last().Value.(int) != node3.Value.(int) {
		t.Errorf("Testing fails: last value error")
	}
}

func TestQueue_delete(t *testing.T) {
	node1 := adt.Node{Value:1}
	node2 := adt.Node{Value:2}
	node3 := adt.Node{Value:3}

	q := MakeQueue()

	q.Add(&node1)
	q.Add(&node2)
	q.Add(&node3)
	q.Delete()

	if q.First().Value.(int) != node2.Value.(int) {
		t.Errorf("Testing fails: first value error")
	}

	if q.Last().Value.(int) != node3.Value.(int) {
		t.Errorf("Testing fails: last value error")
	}
}

func TestQueue_is_empty(t *testing.T) {
	q := MakeQueue()
	if !q.Is_empty() {
		t.Errorf("Testing fails: is_empty error")
	}

	node1 := adt.Node{Value:1}
	q.Add(&node1)
	if q.Is_empty() {
		t.Errorf("Testing fails: is_empty error")
	}
}

