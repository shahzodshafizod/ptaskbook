package dynamic

func SplitList(l1 *List, l2 *List) {
	if l1.Current == nil {
		return
	}
	l2.First = l1.Current
	l2.Last = l1.Last
	l2.Current = l1.Current
	if l1.Current == l1.First {
		l1.First, l1.Last, l1.Current = nil, nil, nil
		return
	}
	l1.Last = l1.Current.Prev
	l1.Last.Next = nil
	l1.Current = l1.First
	l2.First.Prev = nil
}

func AddList(list1, list2 *List) {
	if list1.Current == nil || list2.Current == nil {
		return
	}
	list2.Last.Next = list1.First
	list1.First.Prev = list2.Last
	list2.Last = list1.Last
	list2.Current = list1.First
	list1.First, list1.Last, list1.Current = nil, nil, nil
}

func InsertList(list1, list2 *List) {
	if list1.Current == nil || list2.Current == nil {
		return
	}
	list1.Last.Next = list2.Current
	list1.First.Prev = list2.Current.Prev
	if list2.Current.Prev != nil {
		list2.Current.Prev.Next = list1.First
	}
	list2.Current.Prev = list1.Last
	list2.Current = list1.First
	for list2.First.Prev != nil {
		list2.First = list2.First.Prev
	}
	list1.First, list1.Last, list1.Current = nil, nil, nil
}

func MoreCurrent(list1, list2 *List) {
	if list1.Current == nil {
		return
	}
	list2.InsertAfter(list1.DeleteCurrent())
}
