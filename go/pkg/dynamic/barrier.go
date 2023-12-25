package dynamic

import "github.com/shahzodshafizod/ptaskbook/go/pkg/node"

type BarrierList struct {
	Barrier, Current *node.Node
}

func (b *BarrierList) InsertLast(d int) {
	var newNode = &node.Node{
		Data: d,
		Prev: b.Barrier.Prev,
		Next: b.Barrier,
	}
	b.Barrier.Prev.Next = newNode
	b.Barrier.Prev = newNode
	b.Current = newNode
}

func (b *BarrierList) InsertFirst(d int) {
	var newNode = &node.Node{
		Data: d,
		Next: b.Barrier.Next,
		Prev: b.Barrier,
	}
	b.Barrier.Next.Prev = newNode
	b.Barrier.Next = newNode
	b.Current = newNode
}

func (b *BarrierList) InsertBefore(d int) {
	var newNode = &node.Node{Data: d}
	if b.Current != nil {
		newNode.Prev = b.Current.Prev
		newNode.Next = b.Current
		b.Current.Prev.Next = newNode
		b.Current.Prev = newNode
	} else {
		b.Barrier.Prev, b.Barrier.Next = newNode, newNode
		newNode.Prev, newNode.Next = b.Barrier, b.Barrier
	}
	b.Current = newNode
}

func (b *BarrierList) InsertAfter(d int) {
	var newNode = &node.Node{Data: d}
	if b.Current != nil {
		newNode.Next = b.Current.Next
		newNode.Prev = b.Current
		b.Current.Next.Prev = newNode
		b.Current.Next = newNode
	} else {
		b.Barrier.Next, b.Barrier.Prev = newNode, newNode
		newNode.Next, newNode.Prev = b.Barrier, b.Barrier
	}
	b.Current = newNode
}

func (b *BarrierList) ToFirst() {
	b.Current = b.Barrier.Next
}

func (b *BarrierList) ToLast() {
	b.Current = b.Barrier.Prev
}

func (b *BarrierList) ToNext() {
	b.Current = b.Current.Next
}

func (b *BarrierList) ToPrev() {
	b.Current = b.Current.Prev
}

func (b *BarrierList) SetData(d int) {
	if !b.IsBarrier() {
		b.Current.Data = d
	}
}

func (b *BarrierList) GetData() int {
	return b.Current.Data
}

func (b *BarrierList) IsBarrier() bool {
	return b.Current == b.Barrier
}

func (b *BarrierList) DeleteCurrent() int {
	if b.IsBarrier() {
		return 0
	}
	b.Current.Prev.Next = b.Current.Next
	b.Current.Next.Prev = b.Current.Prev
	var data = b.Current.Data
	if b.Current.Next != b.Barrier {
		b.ToNext()
	} else {
		b.ToPrev()
	}
	return data
}
