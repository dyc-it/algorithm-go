package queue

import (
	"algorithm-go/adt"
	"container/list"
)

type Queue struct {
	list *list.List
}

func MakeQueue() *Queue {
	return &Queue{list:list.New()}
}

func (q *Queue) First() *adt.Node {
	if q.list.Len() == 0 {
		return nil
	} else {
		node := q.list.Front().Value.(*adt.Node)
		return node

	}
}

func (q *Queue) Last() *adt.Node {
	if q.list.Len() == 0 {
		return nil
	} else {
		node := q.list.Back().Value.(*adt.Node)
		return node

	}
}

func (q *Queue) Add(node *adt.Node) {
	q.list.PushBack(node)

}

func (q *Queue) Delete() *adt.Node {
	element := q.list.Front()
	q.list.Remove(element)
	node := element.Value.(*adt.Node)
	return node

}

func (q *Queue) Is_empty() bool {
	return q.list.Len() == 0
}







