package dynamic

import "github.com/shahzodshafizod/ptaskbook/go/pkg/node"

type Queue struct {
	Head, Tail *node.Node
}

func (q *Queue) EnQueue(d int) {
	var newNode = &node.Node{Data: d}
	if q.Tail != nil {
		q.Tail.Next = newNode
	}
	q.Tail = newNode
	if q.Head == nil {
		q.Head = newNode
	}
}

func (q *Queue) DeQueue() int {
	if q.Head == nil {
		return 0
	}
	var data = q.Head.Data
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return data
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}
