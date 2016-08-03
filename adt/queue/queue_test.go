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

	q.add(&node1)
	q.add(&node2)
	q.add(&node3)

	if q.first().Value != node1.Value {
		t.Errorf("Testing fails: first value error")
	}

	if q.last().Value != node3.Value {
		t.Errorf("Testing fails: last value error")
	}
}

func TestQueue_delete(t *testing.T) {
	node1 := adt.Node{Value:1}
	node2 := adt.Node{Value:2}
	node3 := adt.Node{Value:3}

	q := MakeQueue()

	q.add(&node1)
	q.add(&node2)
	q.add(&node3)
	q.delete()

	if q.first().Value != node1.Value {
		t.Errorf("Testing fails: first value error")
	}

	if q.last().Value != node2.Value {
		t.Errorf("Testing fails: last value error")
	}
}
