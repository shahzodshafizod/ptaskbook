package dynamic

import (
	"fmt"

	"github.com/shahzodshafizod/ptaskbook/go/pkg/node"
)

type List struct {
	First, Last, Current *node.Node
}

func (l *List) Println(isDesc bool) {
	var currentData int
	if isDesc {
		for node := l.Last; node != nil; node = node.Prev {
			fmt.Printf("%d\t", node.Data)
			if node == l.Current {
				currentData = node.Data
			}
		}
	} else {
		for node := l.First; node != nil; node = node.Next {
			fmt.Printf("%d\t", node.Data)
			if node == l.Current {
				currentData = node.Data
			}
		}
	}
	fmt.Println(currentData)
}

func (l *List) InsertLast(d int) {
	var newNode = &node.Node{
		Data: d,
		Prev: l.Last,
	}
	if l.Last != nil {
		l.Last.Next = newNode
	}
	l.Last = newNode
	if l.First == nil {
		l.First = newNode
	}
	l.Current = newNode
}

func (l *List) InsertFirst(d int) {
	var newNode = &node.Node{
		Data: d,
		Next: l.First,
	}
	if l.First != nil {
		l.First.Prev = newNode
	}
	l.First = newNode
	if l.Last == nil {
		l.Last = newNode
	}
	l.Current = newNode
}

func (l *List) InsertBefore(d int) {
	var newNode = &node.Node{
		Data: d,
		Next: l.Current,
	}
	if l.Current != nil {
		newNode.Prev = l.Current.Prev
		if l.Current.Prev != nil {
			l.Current.Prev.Next = newNode
		}
		l.Current.Prev = newNode
	} else {
		l.First, l.Last = newNode, newNode
	}
	l.Current = newNode
	if l.First.Prev != nil {
		l.First = l.First.Prev
	}
}

func (l *List) InsertAfter(d int) {
	var newNode = &node.Node{
		Data: d,
		Prev: l.Current,
	}
	if l.Current != nil {
		newNode.Next = l.Current.Next
		if newNode.Next != nil {
			newNode.Next.Prev = newNode
		}
		l.Current.Next = newNode
	} else {
		l.First, l.Last = newNode, newNode
	}
	l.Current = newNode
	if l.Last.Next != nil {
		l.Last = l.Last.Next
	}
}

func (l *List) ToFirst() {
	l.Current = l.First
}

func (l *List) ToLast() {
	l.Current = l.Last
}

func (l *List) ToNext() {
	if l.Current.Next != nil {
		l.Current = l.Current.Next
	}
}

func (l *List) ToPrev() {
	if l.Current.Prev != nil {
		l.Current = l.Current.Prev
	}
}

func (l *List) SetData(d int) {
	if l.Current != nil {
		l.Current.Data = d
	}
}

func (l *List) GetData() int {
	if l.Current != nil {
		return l.Current.Data
	}
	return 0
}

func (l *List) IsLast() bool {
	return l.Current != nil && l.Current == l.Last
}

func (l *List) IsFirst() bool {
	return l.Current != nil && l.Current == l.First
}

func (l *List) DeleteCurrent() int {
	if l.Current == nil {
		return 0
	}
	var data = l.Current.Data
	if l.Current.Prev != nil {
		l.Current.Prev.Next = l.Current.Next
	}
	if l.Current.Next != nil {
		l.Current.Next.Prev = l.Current.Prev
	}
	if l.Current == l.First {
		l.First = l.First.Next
	}
	if l.Current == l.Last {
		l.Last = l.Last.Prev
	}
	if l.Current.Next != nil {
		l.Current = l.Current.Next
	} else {
		l.Current = l.Last
	}
	return data
}
