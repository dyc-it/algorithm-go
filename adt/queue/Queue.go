package queue

import (
	"algorithm-go/adt"
	"container/list"
)

type Queue struct {
	nodes []*adt.Node
	size int
	front int
	rear int
	list *list.List
}

func MakeQueue() *Queue {
	return &Queue{list:list.New()}
}




func (q *Queue) first() *adt.Node {
	if q.list.Len() == 0 {
		return nil
	} else {
		node :=  q.list.Front().Value.(*adt.Node)
		return node
	}
}


func (q *Queue) last() *adt.Node  {
	if q.list.Len() == 0 {
		return nil
	} else {
		node :=  q.list.Back().Value.(*adt.Node)
		return node
	}
}

func (q *Queue) add(node *adt.Node) {
	q.list.PushBack(node)
}

func (q *Queue) delete() *adt.Node  {
	element := q.list.Front()
	q.list.Remove(element)
	node :=  element.Value.(*adt.Node)
	return node
}









