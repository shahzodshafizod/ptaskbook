package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/dynamic"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/node"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type dynamicTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewDynamicTask(pathPrefix string) Task {
	return &dynamicTask{
		dataPath: pathPrefix + "/18dynamic/Dynamic%03d/%03d/%03d",
		name:     "Dynamic",
		count:    80,
	}
}

func (dt *dynamicTask) GetCount() int { return dt.count }

func (dt *dynamicTask) GetName() string { return dt.name }

func (dt *dynamicTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(dt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	dt.scanner = scanner
	dt.checker = checker
	return nil
}

func (dt *dynamicTask) ScannerEOF() bool { return dt.scanner.EOF() }

func (dt *dynamicTask) CheckerEOF() bool { return dt.checker.EOF() }

func (dt *dynamicTask) CleanData() {
	dt.scanner.RemoveFiles()
	dt.checker.RemoveFiles()
}

func (dt *dynamicTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return dt.dynamic1
	case 2:
		return dt.dynamic2
	case 3:
		return dt.dynamic3
	case 4:
		return dt.dynamic4
	case 5:
		return dt.dynamic5
	case 6:
		return dt.dynamic6
	case 7:
		return dt.dynamic7
	case 8:
		return dt.dynamic8
	case 9:
		return dt.dynamic9
	case 10:
		return dt.dynamic10
	case 11:
		return dt.dynamic11
	case 12:
		return dt.dynamic12
	case 13:
		return dt.dynamic13
	case 14:
		return dt.dynamic14
	case 15:
		return dt.dynamic15
	case 16:
		return dt.dynamic16
	case 17:
		return dt.dynamic17
	case 18:
		return dt.dynamic18
	case 19:
		return dt.dynamic19
	case 20:
		return dt.dynamic20
	case 21:
		return dt.dynamic21
	case 22:
		return dt.dynamic22
	case 23:
		return dt.dynamic23
	case 24:
		return dt.dynamic24
	case 25:
		return dt.dynamic25
	case 26:
		return dt.dynamic26
	case 27:
		return dt.dynamic27
	case 28:
		return dt.dynamic28
	case 29:
		return dt.dynamic29
	case 30:
		return dt.dynamic30
	case 31:
		return dt.dynamic31
	case 32:
		return dt.dynamic32
	case 33:
		return dt.dynamic33
	case 34:
		return dt.dynamic34
	case 35:
		return dt.dynamic35
	case 36:
		return dt.dynamic36
	case 37:
		return dt.dynamic37
	case 38:
		return dt.dynamic38
	case 39:
		return dt.dynamic39
	case 40:
		return dt.dynamic40
	case 41:
		return dt.dynamic41
	case 42:
		return dt.dynamic42
	case 43:
		return dt.dynamic43
	case 44:
		return dt.dynamic44
	case 45:
		return dt.dynamic45
	case 46:
		return dt.dynamic46
	case 47:
		return dt.dynamic47
	case 48:
		return dt.dynamic48
	case 49:
		return dt.dynamic49
	case 50:
		return dt.dynamic50
	case 51:
		return dt.dynamic51
	case 52:
		return dt.dynamic52
	case 53:
		return dt.dynamic53
	case 54:
		return dt.dynamic54
	case 55:
		return dt.dynamic55
	case 56:
		return dt.dynamic56
	case 57:
		return dt.dynamic57
	case 58:
		return dt.dynamic58
	case 59:
		return dt.dynamic59
	case 60:
		return dt.dynamic60
	case 61:
		return dt.dynamic61
	case 62:
		return dt.dynamic62
	case 63:
		return dt.dynamic63
	case 64:
		return dt.dynamic64
	case 65:
		return dt.dynamic65
	case 66:
		return dt.dynamic66
	case 67:
		return dt.dynamic67
	case 68:
		return dt.dynamic68
	case 69:
		return dt.dynamic69
	case 70:
		return dt.dynamic70
	case 71:
		return dt.dynamic71
	case 72:
		return dt.dynamic72
	case 73:
		return dt.dynamic73
	case 74:
		return dt.dynamic74
	case 75:
		return dt.dynamic75
	case 76:
		return dt.dynamic76
	case 77:
		return dt.dynamic77
	case 78:
		return dt.dynamic78
	case 79:
		return dt.dynamic79
	case 80:
		return dt.dynamic80
	default:
		return nil
	}
}

func (dt *dynamicTask) dynamic1() bool {
	var p1 = dt.scanner.NextStack()
	if !dt.checker.CompareInt(p1.Data, p1.Next.Data) {
		return false
	}
	var p2 = p1.Next
	return dt.checker.CompareInt(p2.Data)
}

func (dt *dynamicTask) dynamic2() bool {
	var top = dt.scanner.NextStack()
	var count = 0
	for i := top; i != nil; i = i.Next {
		count++
		if !dt.checker.CompareInt(i.Data) {
			return false
		}
	}
	return dt.checker.CompareInt(count) && dt.checker.CompareStack(top)
}

func (dt *dynamicTask) dynamic3() bool {
	var d = dt.scanner.NextInt()
	var top = dt.scanner.NextStack()
	var newNode = &node.Node{
		Data: d,
		Next: top,
	}
	top = newNode
	return dt.checker.CompareStack(top)
}

func (dt *dynamicTask) dynamic4() bool {
	var n = dt.scanner.NextInt()
	var data int
	var top, newNode *node.Node
	for index := 0; index < n; index++ {
		data = dt.scanner.NextInt()
		newNode = &node.Node{
			Data: data,
			Next: top,
		}
		top = newNode
	}
	return dt.checker.CompareStack(top)
}

func (dt *dynamicTask) dynamic5() bool {
	var p1 = dt.scanner.NextStack()
	var d = p1.Data
	if !dt.checker.CompareInt(d) {
		return false
	}
	var p2 = p1.Next
	return dt.checker.CompareStack(p2)
}

func (dt *dynamicTask) dynamic6() bool {
	var p1 = dt.scanner.NextStack()
	for index := 0; index < 9; index++ {
		if !dt.checker.CompareInt(p1.Data) {
			return false
		}
		p1 = p1.Next
	}
	var p2 = p1
	return dt.checker.CompareStack(p2)
}

func (dt *dynamicTask) dynamic7() bool {
	var p1 = dt.scanner.NextStack()
	var count = 0
	for p1 != nil {
		if !dt.checker.CompareInt(p1.Data) {
			return false
		}
		p1 = p1.Next
		count++
	}
	return dt.checker.CompareInt(count)
}

func (dt *dynamicTask) dynamic8() bool {
	var p1 = dt.scanner.NextStack()
	var p2 = dt.scanner.NextStack()
	var tempNode *node.Node
	for p1 != nil {
		tempNode = p1
		p1 = p1.Next
		tempNode.Next = p2
		p2 = tempNode
	}
	return dt.checker.CompareStack(p2)
}

func (dt *dynamicTask) dynamic9() bool {
	var p1 = dt.scanner.NextStack()
	var p2 = dt.scanner.NextStack()
	var tempNode *node.Node
	for p1 != nil && p1.Data%2 != 0 {
		tempNode = p1
		p1 = p1.Next
		tempNode.Next = p2
		p2 = tempNode
	}
	return dt.checker.CompareStack(p1) && dt.checker.CompareStack(p2)
}

func (dt *dynamicTask) dynamic10() bool {
	var p1 = dt.scanner.NextStack()
	var evens, unevens, tempNode *node.Node
	for p1 != nil {
		tempNode = p1
		p1 = p1.Next
		if tempNode.Data%2 == 0 {
			tempNode.Next = evens
			evens = tempNode
		} else {
			tempNode.Next = unevens
			unevens = tempNode
		}
	}
	return dt.checker.CompareStack(evens) && dt.checker.CompareStack(unevens)
}

func (dt *dynamicTask) dynamic11() bool {
	var p1 = dt.scanner.NextStack()
	var ts = &dynamic.Stack{Top: p1}
	var n = dt.scanner.NextInt()
	var data int
	for index := 0; index < n; index++ {
		data = dt.scanner.NextInt()
		ts.Push(data)
	}
	return dt.checker.CompareStack(ts.Top)
}

func (dt *dynamicTask) dynamic12() bool {
	var p1 = dt.scanner.NextStack()
	var ts = &dynamic.Stack{Top: p1}
	for index := 0; index < 5; index++ {
		if !dt.checker.CompareInt(ts.Pop()) {
			return false
		}
	}
	return dt.checker.CompareStack(ts.Top)
}

func (dt *dynamicTask) dynamic13() bool {
	var p1 = dt.scanner.NextStack()
	var ts = &dynamic.Stack{Top: p1}
	for index := 0; index < 5; index++ {
		if !ts.IsEmpty() {
			if !dt.checker.CompareInt(ts.Pop()) {
				return false
			}
		}
	}
	if !dt.checker.CompareBool(ts.IsEmpty()) {
		return false
	}
	if !ts.IsEmpty() {
		return dt.checker.CompareStack(ts.Top)
	}
	return true
}

func (dt *dynamicTask) dynamic14() bool {
	var data int
	var p1, p2, newNode *node.Node
	for index := 0; index < 10; index++ {
		data = dt.scanner.NextInt()
		newNode = &node.Node{Data: data}
		if p2 != nil {
			p2.Next = newNode
		}
		p2 = newNode
		if p1 == nil {
			p1 = newNode
		}
	}
	return dt.checker.CompareQueue(p1, p2)
}

func (dt *dynamicTask) dynamic15() bool {
	var data int
	var (
		newNode                *node.Node
		firstHead, firstTail   *node.Node
		secondHead, secondTail *node.Node
		head, tail             **node.Node
	)
	for index := 1; index <= 10; index++ {
		data = dt.scanner.NextInt()
		newNode = &node.Node{Data: data}
		if index%2 != 0 {
			head, tail = &firstHead, &firstTail
		} else {
			head, tail = &secondHead, &secondTail
		}
		if *tail != nil {
			(*tail).Next = newNode
		}
		*tail = newNode
		if *head == nil {
			*head = newNode
		}
	}
	return dt.checker.CompareQueue(firstHead, firstTail) &&
		dt.checker.CompareQueue(secondHead, secondTail)
}

func (dt *dynamicTask) dynamic16() bool {
	var data int
	var (
		newNode                *node.Node
		firstHead, firstTail   *node.Node
		secondHead, secondTail *node.Node
		head, tail             **node.Node
	)
	for index := 0; index < 10; index++ {
		data = dt.scanner.NextInt()
		newNode = &node.Node{Data: data}
		if data%2 != 0 {
			head, tail = &firstHead, &firstTail
		} else {
			head, tail = &secondHead, &secondTail
		}
		if *tail != nil {
			(*tail).Next = newNode
		}
		*tail = newNode
		if *head == nil {
			*head = newNode
		}
	}
	return dt.checker.CompareQueue(firstHead, firstTail) &&
		dt.checker.CompareQueue(secondHead, secondTail)
}

func (dt *dynamicTask) dynamic17() bool {
	var d = dt.scanner.NextInt()
	var p1, p2 = dt.scanner.NextQueue()
	var newNode = &node.Node{Data: d}
	if p2 != nil {
		p2.Next = newNode
	}
	p2 = newNode
	p2.Next = nil
	if p1 == nil {
		p1 = newNode
	}
	return dt.checker.CompareQueue(p1, p2)
}

func (dt *dynamicTask) dynamic18() bool {
	var d = dt.scanner.NextInt()
	var p1, p2 = dt.scanner.NextQueue()
	var newNode = &node.Node{Data: d}
	if p2 != nil {
		p2.Next = newNode
	}
	p2 = newNode
	p2.Next = nil
	if p1 == nil {
		p1 = newNode
	}
	d = p1.Data
	p1 = p1.Next
	if !dt.checker.CompareInt(d) {
		return false
	}
	return dt.checker.CompareQueue(p1, p2)
}

func (dt *dynamicTask) dynamic19() bool {
	var n = dt.scanner.NextInt()
	var p1, p2 = dt.scanner.NextQueue()
	for index := 0; index < n && p1 != nil; index++ {
		if !dt.checker.CompareInt(p1.Data) {
			return false
		}
		p1 = p1.Next
		if p1 == nil {
			p2 = nil
		}
	}
	return dt.checker.CompareQueue(p1, p2)
}

func (dt *dynamicTask) dynamic20() bool {
	var p1, p2 = dt.scanner.NextQueue()
	for p1 != nil && p1.Data%2 != 0 {
		if !dt.checker.CompareInt(p1.Data) {
			return false
		}
		p1 = p1.Next
		if p1 == nil {
			p2 = nil
		}
	}
	return dt.checker.CompareQueue(p1, p2)
}

func (dt *dynamicTask) dynamic21() bool {
	var p1, _ = dt.scanner.NextQueue()
	var p3, p4 = dt.scanner.NextQueue()
	var tempNode *node.Node
	for p1 != nil {
		tempNode = p1
		p1 = p1.Next
		if p4 != nil {
			p4.Next = tempNode
		}
		p4 = tempNode
		p4.Next = nil
		if p3 == nil {
			p3 = tempNode
		}
	}
	return dt.checker.CompareQueue(p3, p4)
}

func (dt *dynamicTask) dynamic22() bool {
	var n = dt.scanner.NextInt()
	var p1, p2 = dt.scanner.NextQueue()
	var p3, p4 = dt.scanner.NextQueue()
	var tempNode *node.Node
	for index := 0; index < n && p1 != nil; index++ {
		tempNode = p1
		p1 = p1.Next
		if p4 != nil {
			p4.Next = tempNode
		}
		p4 = tempNode
		p4.Next = nil
		if p3 == nil {
			p3 = tempNode
		}
	}
	if p1 == nil {
		p2 = nil
	}
	return dt.checker.CompareQueue(p1, p2) && dt.checker.CompareQueue(p3, p4)
}

func (dt *dynamicTask) dynamic23() bool {
	var p1, p2 = dt.scanner.NextQueue()
	var p3, p4 = dt.scanner.NextQueue()
	var tempNode *node.Node
	for p1 != nil && p1.Data%2 != 0 {
		tempNode = p1
		p1 = p1.Next
		if p4 != nil {
			p4.Next = tempNode
		}
		p4 = tempNode
		p4.Next = nil
		if p3 == nil {
			p3 = tempNode
		}
	}
	if p1 == nil {
		p2 = nil
	}
	return dt.checker.CompareQueue(p1, p2) && dt.checker.CompareQueue(p3, p4)
}

func (dt *dynamicTask) dynamic24() bool {
	var p1, _ = dt.scanner.NextQueue()
	var p3, _ = dt.scanner.NextQueue()
	var head, tail, tempNode *node.Node
	for p1 != nil && p3 != nil {
		tempNode = p1
		p1 = p1.Next
		tempNode.Next = nil
		if tail != nil {
			tail.Next = tempNode
		}
		tail = tempNode
		if head == nil {
			head = tempNode
		}

		tempNode = p3
		p3 = p3.Next
		tempNode.Next = nil
		tail.Next = tempNode
		tail = tempNode
	}
	return dt.checker.CompareQueue(head, tail)
}

func (dt *dynamicTask) dynamic25() bool {
	var p1, _ = dt.scanner.NextQueue()
	var p3, _ = dt.scanner.NextQueue()
	var head, tail, tempNode *node.Node
	for p1 != nil || p3 != nil {
		if p1 == nil {
			tempNode = p3
			p3 = p3.Next
		} else if p3 == nil {
			tempNode = p1
			p1 = p1.Next
		} else if p1.Data < p3.Data {
			tempNode = p1
			p1 = p1.Next
		} else {
			tempNode = p3
			p3 = p3.Next
		}
		tempNode.Next = nil
		if tail != nil {
			tail.Next = tempNode
		}
		tail = tempNode
		if head == nil {
			head = tempNode
		}
	}
	return dt.checker.CompareQueue(head, tail)
}

func (dt *dynamicTask) dynamic26() bool {
	var p1, p2 = dt.scanner.NextQueue()
	var tq = &dynamic.Queue{Head: p1, Tail: p2}
	var n = dt.scanner.NextInt()
	var data int
	for index := 0; index < n; index++ {
		data = dt.scanner.NextInt()
		tq.EnQueue(data)
	}
	return dt.checker.CompareQueue(tq.Head, tq.Tail)
}

func (dt *dynamicTask) dynamic27() bool {
	var p1, p2 = dt.scanner.NextQueue()
	var tq = dynamic.Queue{Head: p1, Tail: p2}
	for index := 0; index < 5; index++ {
		if !dt.checker.CompareInt(tq.DeQueue()) {
			return false
		}
	}
	return dt.checker.CompareQueue(tq.Head, tq.Tail)
}

func (dt *dynamicTask) dynamic28() bool {
	var p1, p2 = dt.scanner.NextQueue()
	var tq = dynamic.Queue{Head: p1, Tail: p2}
	for index := 0; index < 5 && !tq.IsEmpty(); index++ {
		if !dt.checker.CompareInt(tq.DeQueue()) {
			return false
		}
	}
	return dt.checker.CompareBool(tq.IsEmpty()) && dt.checker.CompareQueue(tq.Head, tq.Tail)
}

func (dt *dynamicTask) dynamic29() bool {
	var _, _, p2 = dt.scanner.NextList()
	if !dt.checker.CompareInt(p2.Prev.Data, p2.Next.Data) {
		return false
	}
	return dt.checker.CompareList(p2.Prev, p2.Next, p2)
}

func (dt *dynamicTask) dynamic30() bool {
	var p1, _ = dt.scanner.NextQueue()
	var nextNode = p1
	for nextNode != nil {
		p1 = nextNode
		nextNode = nextNode.Next
		if nextNode != nil {
			nextNode.Prev = p1
		}
	}
	return dt.checker.CompareList(nil, p1, nil)
}

func (dt *dynamicTask) dynamic31() bool {
	var _, _, p0 = dt.scanner.NextList()
	var p1 = p0
	for p1 != nil && p1.Prev != nil {
		p1 = p1.Prev
	}
	var p2, nextNode = p1, p1
	var n = 0
	for nextNode != nil {
		n++
		p2 = nextNode
		nextNode = nextNode.Next
	}
	return dt.checker.CompareInt(n) && dt.checker.CompareList(p1, p2, nil)
}

func (dt *dynamicTask) dynamic32() bool {
	var d1 = dt.scanner.NextInt()
	var d2 = dt.scanner.NextInt()
	var _, _, p0 = dt.scanner.NextList()
	var first, last = p0, p0
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	for last != nil && last.Next != nil {
		last = last.Next
	}
	if first != nil {
		first.Prev = &node.Node{
			Data: d1,
			Next: first,
		}
		first = first.Prev
	} else {
		first = &node.Node{Data: d1}
	}
	if last != nil {
		last.Next = &node.Node{
			Data: d2,
			Prev: last,
		}
		last = last.Next
	} else {
		last = &node.Node{Data: d2}
	}
	return dt.checker.CompareList(first, last, nil)
}

func (dt *dynamicTask) dynamic33() bool {
	var d = dt.scanner.NextInt()
	var _, _, p0 = dt.scanner.NextList()
	var newNode = &node.Node{
		Data: d,
		Next: p0,
	}
	if p0 != nil {
		if p0.Prev != nil {
			p0.Prev.Next = newNode
		}
		newNode.Prev = p0.Prev
		p0.Prev = newNode
	}
	p0 = newNode
	return dt.checker.CompareList(nil, nil, p0)
}

func (dt *dynamicTask) dynamic34() bool {
	var d = dt.scanner.NextInt()
	var _, _, p0 = dt.scanner.NextList()
	var newNode = &node.Node{
		Data: d,
		Prev: p0,
	}
	if p0 != nil {
		if p0.Next != nil {
			p0.Next.Prev = newNode
		}
		newNode.Next = p0.Next
		p0.Next = newNode
	}
	p0 = newNode
	return dt.checker.CompareList(nil, nil, p0)
}

func (dt *dynamicTask) dynamic35() bool {
	var p1, p2, _ = dt.scanner.NextList()
	if p1 != nil {
		p1.Prev = &node.Node{
			Data: p1.Data,
			Next: p1,
		}
		p1 = p1.Prev
	}
	if p2 != nil {
		var newNode = &node.Node{
			Data: p2.Data,
			Next: p2,
		}
		if p2.Prev != nil {
			p2.Prev.Next = newNode
			newNode.Prev = p2.Prev
		}
		p2.Prev = newNode
	}
	return dt.checker.CompareList(p1, p2, nil)
}

func (dt *dynamicTask) dynamic36() bool {
	var p1, p2, _ = dt.scanner.NextList()
	if p1 != nil {
		var newNode = &node.Node{
			Data: p1.Data,
			Prev: p1,
		}
		if p1.Next != nil {
			p1.Next.Prev = newNode
			newNode.Next = p1.Next
		}
		p1.Next = newNode
	}
	if p2 != nil {
		p2.Next = &node.Node{
			Data: p2.Data,
			Prev: p2,
		}
		p2 = p2.Next
	}
	return dt.checker.CompareList(p1, p2, nil)
}

func (dt *dynamicTask) dynamic37() bool {
	var p1, _, _ = dt.scanner.NextList()
	var index = 0
	var newNode *node.Node
	for p := p1; p != nil; p = p.Next {
		index++
		if index%2 != 0 {
			newNode = &node.Node{
				Data: p.Data,
				Next: p,
			}
			if p.Prev != nil {
				p.Prev.Next = newNode
				newNode.Prev = p.Prev
			}
			p.Prev = newNode
		}
	}
	for p1 != nil && p1.Prev != nil {
		p1 = p1.Prev
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic38() bool {
	var p1, _, _ = dt.scanner.NextList()
	var index = 0
	var newNode *node.Node
	for p := p1; p != nil; p = p.Next {
		index++
		if index%2 != 0 {
			newNode = &node.Node{
				Data: p.Data,
				Prev: p,
			}
			if p.Next != nil {
				p.Next.Prev = newNode
				newNode.Next = p.Next
			}
			p.Next = newNode
			p = p.Next
		}
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic39() bool {
	var p1, _, _ = dt.scanner.NextList()
	var newNode *node.Node
	for p := p1; p != nil; p = p.Next {
		if p.Data%2 != 0 {
			newNode = &node.Node{
				Data: p.Data,
				Next: p,
			}
			if p.Prev != nil {
				p.Prev.Next = newNode
				newNode.Prev = p.Prev
			}
			p.Prev = newNode
		}
	}
	for p1 != nil && p1.Prev != nil {
		p1 = p1.Prev
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic40() bool {
	var p1, _, _ = dt.scanner.NextList()
	var newNode *node.Node
	for p := p1; p != nil; p = p.Next {
		if p.Data%2 != 0 {
			newNode = &node.Node{
				Data: p.Data,
				Prev: p,
			}
			if p.Next != nil {
				p.Next.Prev = newNode
				newNode.Next = p.Next
			}
			p.Next = newNode
			p = p.Next
		}
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic41() bool {
	var _, _, p0 = dt.scanner.NextList()
	if p0.Prev != nil {
		p0.Prev.Next = p0.Next
	}
	if p0.Next != nil {
		p0.Next.Prev = p0.Prev
	}
	var first = p0.Prev
	if first == nil {
		first = p0.Next
	}
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	var orderNumber = 0
	for index, node := 1, first; node != nil; index, node = index+1, node.Next {
		if node == p0.Prev {
			orderNumber = index
			break
		}
	}
	if !dt.checker.CompareInt(orderNumber) {
		return false
	}
	orderNumber = 0
	for index, node := 1, first; node != nil; index, node = index+1, node.Next {
		if node == p0.Next {
			orderNumber = index
			break
		}
	}
	return dt.checker.CompareInt(orderNumber) && dt.checker.CompareList(first, nil, nil)
}

func (dt *dynamicTask) dynamic42() bool {
	var p1, _, _ = dt.scanner.NextList()
	var index = 0
	for p := p1; p != nil; p = p.Next {
		index++
		if index%2 != 0 {
			if p.Prev != nil {
				p.Prev.Next = p.Next
			}
			if p.Next != nil {
				p.Next.Prev = p.Prev
			}
			if p == p1 {
				p1 = p1.Next
			}
		}
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic43() bool {
	var p1, _, _ = dt.scanner.NextList()
	for p := p1; p != nil; p = p.Next {
		if p.Data%2 != 0 {
			if p.Prev != nil {
				p.Prev.Next = p.Next
			}
			if p.Next != nil {
				p.Next.Prev = p.Prev
			}
			if p == p1 {
				p1 = p1.Next
			}
		}
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic44() bool {
	var _, _, p0 = dt.scanner.NextList()
	for p0 != nil && p0.Next != nil {
		p0.Next.Prev = p0.Prev
		if p0.Prev != nil {
			p0.Prev.Next = p0.Next
		}
		if p0.Next.Next != nil {
			p0.Next.Next.Prev = p0
		}
		p0.Prev = p0.Next
		p0.Next = p0.Next.Next
		p0.Prev.Next = p0
	}
	var first = p0
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	return dt.checker.CompareList(first, p0, nil)
}

func (dt *dynamicTask) dynamic45() bool {
	var _, _, p0 = dt.scanner.NextList()
	for p0 != nil && p0.Prev != nil {
		p0.Prev.Next = p0.Next
		if p0.Next != nil {
			p0.Next.Prev = p0.Prev
		}
		if p0.Prev.Prev != nil {
			p0.Prev.Prev.Next = p0
		}
		p0.Next = p0.Prev
		p0.Prev = p0.Prev.Prev
		p0.Next.Prev = p0
	}
	var last = p0
	for last != nil && last.Next != nil {
		last = last.Next
	}
	return dt.checker.CompareList(p0, last, nil)
}

func (dt *dynamicTask) dynamic46() bool {
	var k = dt.scanner.NextInt()
	var _, _, p0 = dt.scanner.NextList()
	var index = 0
	for p0 != nil && p0.Next != nil {
		index++
		if index > k {
			break
		}
		p0.Next.Prev = p0.Prev
		if p0.Prev != nil {
			p0.Prev.Next = p0.Next
		}
		if p0.Next.Next != nil {
			p0.Next.Next.Prev = p0
		}
		p0.Prev = p0.Next
		p0.Next = p0.Next.Next
		p0.Prev.Next = p0
	}
	var first = p0
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	var last = p0
	for last != nil && last.Next != nil {
		last = last.Next
	}
	return dt.checker.CompareList(first, last, nil)
}

func (dt *dynamicTask) dynamic47() bool {
	var k = dt.scanner.NextInt()
	var _, _, p0 = dt.scanner.NextList()
	var index = 0
	for p0 != nil && p0.Prev != nil {
		index++
		if index > k {
			break
		}
		p0.Prev.Next = p0.Next
		if p0.Next != nil {
			p0.Next.Prev = p0.Prev
		}
		if p0.Prev.Prev != nil {
			p0.Prev.Prev.Next = p0
		}
		p0.Next = p0.Prev
		p0.Prev = p0.Prev.Prev
		p0.Next.Prev = p0
	}
	var first = p0
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	var last = p0
	for last != nil && last.Next != nil {
		last = last.Next
	}
	return dt.checker.CompareList(first, last, nil)
}

func (dt *dynamicTask) dynamic48() bool {
	var x = dt.scanner.NextInt()
	var y = dt.scanner.NextInt()
	var p, _, _ = dt.scanner.NextList()
	var px, py *node.Node = nil, nil
	for index := 1; (p != nil) && (px == nil || py == nil); index, p = index+1, p.Next {
		switch index {
		case x:
			px = p
		case y:
			py = p
		}
	}
	// move px after py: n steps;
	var stepsCount = 0
	for px != nil && px.Next != nil {
		stepsCount++
		px.Next.Prev = px.Prev
		if px.Prev != nil {
			px.Prev.Next = px.Next
		}
		if px.Next.Next != nil {
			px.Next.Next.Prev = px
		}
		px.Prev = px.Next
		px.Next = px.Next.Next
		px.Prev.Next = px
		if px.Prev == py {
			break
		}
	}
	// move py to the beginning: n-1 steps;
	for py != nil && py.Prev != nil {
		stepsCount--
		if stepsCount <= 0 {
			break
		}
		py.Prev.Next = py.Next
		if py.Next != nil {
			py.Next.Prev = py.Prev
		}
		if py.Prev.Prev != nil {
			py.Prev.Prev.Next = py
		}
		py.Next = py.Prev
		py.Prev = py.Prev.Prev
		py.Next.Prev = py
	}
	for py != nil && py.Prev != nil {
		py = py.Prev
	}
	for px != nil && px.Next != nil {
		px = px.Next
	}
	return dt.checker.CompareList(py, px, nil)
}

func (dt *dynamicTask) dynamic49() bool {
	var p1, _, _ = dt.scanner.NextList()
	var node *node.Node
	var count = 0
	for node = p1; node != nil; node = node.Next {
		count++
	}
	var index = 0
	for p1 != nil && p1.Next != nil {
		node = p1
		p1 = p1.Next
		index++
		if index > count {
			break
		}
		if index%2 != 0 {
			for node.Next != nil {
				node.Next.Prev = node.Prev
				if node.Prev != nil {
					node.Prev.Next = node.Next
				}
				if node.Next.Next != nil {
					node.Next.Next.Prev = node
				}
				node.Prev = node.Next
				node.Next = node.Next.Next
				node.Prev.Next = node
			}
		}
	}
	for p1 != nil && p1.Prev != nil {
		p1 = p1.Prev
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic50() bool {
	var p1, _, _ = dt.scanner.NextList()
	var node *node.Node
	var count = 0
	for node = p1; node != nil; node = node.Next {
		count++
	}
	var index = 0
	for p1 != nil && p1.Next != nil {
		node = p1
		p1 = p1.Next
		index++
		if index > count {
			break
		}
		if node.Data%2 != 0 {
			for node.Next != nil {
				node.Next.Prev = node.Prev
				if node.Prev != nil {
					node.Prev.Next = node.Next
				}
				if node.Next.Next != nil {
					node.Next.Next.Prev = node
				}
				node.Prev = node.Next
				node.Next = node.Next.Next
				node.Prev.Next = node
			}
		}
	}
	for p1 != nil && p1.Prev != nil {
		p1 = p1.Prev
	}
	return dt.checker.CompareList(p1, nil, nil)
}

func (dt *dynamicTask) dynamic51() bool {
	var p1, _, _ = dt.scanner.NextList()
	var _, _, p0 = dt.scanner.NextList()
	var tempNode *node.Node
	for p1 != nil {
		tempNode = p1
		p1 = p1.Next
		tempNode.Prev, tempNode.Next = nil, nil
		if p0 == nil {
			p0 = tempNode
		} else {
			if p0.Prev != nil {
				p0.Prev.Next = tempNode
			}
			tempNode.Prev = p0.Prev
			tempNode.Next = p0
			p0.Prev = tempNode
		}
	}
	var first = p0
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	var last = p0
	for last != nil && last.Next != nil {
		last = last.Next
	}
	return dt.checker.CompareList(first, last, nil)
}

func (dt *dynamicTask) dynamic52() bool {
	var p1, _, _ = dt.scanner.NextList()
	var _, _, p0 = dt.scanner.NextList()
	var tempNode *node.Node
	for p1 != nil {
		tempNode = p1
		p1 = p1.Next
		tempNode.Prev, tempNode.Next = nil, nil
		if p0 != nil {
			if p0.Next != nil {
				p0.Next.Prev = tempNode
			}
			tempNode.Next = p0.Next
			tempNode.Prev = p0
			p0.Next = tempNode
		}
		p0 = tempNode
	}
	var first = p0
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	var last = p0
	for last != nil && last.Next != nil {
		last = last.Next
	}
	return dt.checker.CompareList(first, last, nil)
}

func (dt *dynamicTask) dynamic53() bool {
	var x = dt.scanner.NextInt()
	var y = dt.scanner.NextInt()
	var p1, _, _ = dt.scanner.NextList()
	var px, py *node.Node
	for index, node := 1, p1; node != nil; index, node = index+1, node.Next {
		if index == x {
			px = node
			break
		}
	}
	for index, node := 1, p1; node != nil; index, node = index+1, node.Next {
		if index == y {
			py = node
			break
		}
	}
	var first, last = px.Prev, py.Next
	px.Prev, py.Next = nil, nil
	if first != nil {
		first.Next = last
	}
	if last != nil {
		last.Prev = first
	}
	if first == nil {
		first = last
	}
	for first != nil && first.Prev != nil {
		first = first.Prev
	}
	if last == nil {
		last = first
	}
	for last != nil && last.Next != nil {
		last = last.Next
	}
	return dt.checker.CompareList(first, last, nil) && dt.checker.CompareList(px, py, nil)
}

func (dt *dynamicTask) dynamic54() bool {
	var x = dt.scanner.NextInt()
	var y = dt.scanner.NextInt()
	var p1, _, _ = dt.scanner.NextList()
	var px, py *node.Node
	for index, node := 1, p1; node != nil; index, node = index+1, node.Next {
		if index == x {
			px = node
			break
		}
	}
	for index, node := 1, p1; node != nil; index, node = index+1, node.Next {
		if index == y {
			py = node
			break
		}
	}
	var first, last = px.Next, py.Prev
	if first.Prev == last {
		first, last = nil, nil
	}
	px.Next, py.Prev = py, px
	if first != nil {
		first.Prev = nil
	}
	if last != nil {
		last.Next = nil
	}
	if first == nil {
		first = last
		for first != nil && first.Prev != nil {
			first = first.Prev
		}
	}
	if last == nil {
		last = first
		for last != nil && last.Next != nil {
			last = last.Next
		}
	}
	for px != nil && px.Prev != nil {
		px = px.Prev
	}
	for py != nil && py.Next != nil {
		py = py.Next
	}
	return dt.checker.CompareList(px, py, nil) && dt.checker.CompareList(first, last, nil)
}

func (dt *dynamicTask) dynamic55() bool {
	var p1, _, _ = dt.scanner.NextList()
	var p2 = p1
	for p2 != nil && p2.Next != nil {
		p2 = p2.Next
	}
	if p1 != nil {
		p1.Prev = p2
	}
	if p2 != nil {
		p2.Next = p1
	}
	return dt.checker.CompareList(nil, p2, nil)
}

func (dt *dynamicTask) dynamic56() bool {
	var p1, p2, _ = dt.scanner.NextList()
	var elementsCount = 0
	for p := p1; p != nil; p = p.Next {
		elementsCount++
	}
	var p3, p4 *node.Node
	p3 = p1
	for index := 1; index < elementsCount/2; index++ {
		p3 = p3.Next
	}
	if p3 != nil {
		p4 = p3.Next
	}
	if p1 != nil {
		p1.Prev = p3
	}
	if p2 != nil {
		p2.Next = p4
	}
	return dt.checker.CompareList(nil, p3, nil) && dt.checker.CompareList(p4, nil, nil)
}

func (dt *dynamicTask) dynamic57() bool {
	var k = dt.scanner.NextInt()
	var p1, p2, _ = dt.scanner.NextList()
	var tempNode *node.Node
	for index := 0; index < k && p1 != nil && p2 != nil; index++ {
		tempNode = p2
		p2 = p2.Prev
		p2.Next = nil
		tempNode.Prev = nil
		tempNode.Next = p1
		p1.Prev = tempNode
		p1 = tempNode
		tempNode = nil
	}
	return dt.checker.CompareList(p1, p2, nil)
}

func (dt *dynamicTask) dynamic58() bool {
	var k = dt.scanner.NextInt()
	var p1, p2, _ = dt.scanner.NextList()
	var tempNode *node.Node
	for index := 0; index < k && p1 != nil && p2 != nil; index++ {
		tempNode = p1
		p1 = p1.Next
		p1.Prev = nil
		tempNode.Next = nil
		tempNode.Prev = p2
		p2.Next = tempNode
		p2 = tempNode
		tempNode = nil
	}
	return dt.checker.CompareList(p1, p2, nil)
}

func (dt *dynamicTask) dynamic59() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var tl = &dynamic.List{
		First:   p1,
		Last:    p2,
		Current: p3,
	}
	var n = dt.scanner.NextInt()
	var data int
	for index := 0; index < n; index++ {
		data = dt.scanner.NextInt()
		tl.InsertLast(data)
	}
	return dt.checker.CompareList(tl.First, tl.Last, tl.Current)
}

func (dt *dynamicTask) dynamic60() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var tl = &dynamic.List{
		First:   p1,
		Last:    p2,
		Current: p3,
	}
	var n = dt.scanner.NextInt()
	var data int
	for index := 0; index < n; index++ {
		data = dt.scanner.NextInt()
		tl.InsertFirst(data)
	}
	return dt.checker.CompareList(tl.First, tl.Last, tl.Current)
}

func (dt *dynamicTask) dynamic61() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var data int
	var tl = &dynamic.List{
		First:   p1,
		Last:    p2,
		Current: p3,
	}
	for index := 0; index < 5; index++ {
		data = dt.scanner.NextInt()
		tl.InsertBefore(data)
	}
	return dt.checker.CompareList(tl.First, tl.Last, tl.Current)
}

func (dt *dynamicTask) dynamic62() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var data int
	var tl = &dynamic.List{
		First:   p1,
		Last:    p2,
		Current: p3,
	}
	for index := 0; index < 5; index++ {
		data = dt.scanner.NextInt()
		tl.InsertAfter(data)
	}
	return dt.checker.CompareList(tl.First, tl.Last, tl.Current)
}

func (dt *dynamicTask) dynamic63() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var tl = &dynamic.List{
		First:   p1,
		Last:    p2,
		Current: p3,
	}
	tl.ToFirst()
	var count = 0
	for {
		count++
		if count%2 != 0 {
			tl.SetData(0)
		}
		if tl.IsLast() {
			break
		}
		tl.ToNext()
	}
	return dt.checker.CompareInt(count) && dt.checker.CompareList(tl.First, tl.Last, tl.Current)
}

func (dt *dynamicTask) dynamic64() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var tl = &dynamic.List{
		First:   p1,
		Last:    p2,
		Current: p3,
	}
	tl.ToLast()
	var count = 0
	var data int
	for tl.Current != nil {
		count++
		data = tl.GetData()
		if data%2 == 0 {
			if !dt.checker.CompareInt(data) {
				return false
			}
		}
		if tl.IsFirst() {
			break
		}
		tl.ToPrev()
	}
	return dt.checker.CompareInt(count)
}

func (dt *dynamicTask) dynamic65() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var tl = &dynamic.List{First: p1, Last: p2, Current: p3}
	for index := 0; index < 5 && tl.Current != nil; index++ {
		if !dt.checker.CompareInt(tl.DeleteCurrent()) {
			return false
		}
	}
	return dt.checker.CompareList(tl.First, tl.Last, tl.Current)
}

func (dt *dynamicTask) dynamic66() bool {
	var p1, p2, p3 = dt.scanner.NextList()
	var tl1 = &dynamic.List{First: p1, Last: p2, Current: p3}
	var tl2 = &dynamic.List{}
	dynamic.SplitList(tl1, tl2)
	return dt.checker.CompareList(tl1.First, tl1.Last, tl1.Current) &&
		dt.checker.CompareList(tl2.First, tl2.Last, tl2.Current)
}

func (dt *dynamicTask) dynamic67() bool {
	var list1, list2 = &dynamic.List{}, &dynamic.List{}
	list1.First, list1.Last, list1.Current = dt.scanner.NextList()
	list2.First, list2.Last, list2.Current = dt.scanner.NextList()
	dynamic.AddList(list1, list2)
	return dt.checker.CompareList(list2.First, list2.Last, list2.Current)
}

func (dt *dynamicTask) dynamic68() bool {
	var list1, list2 = &dynamic.List{}, &dynamic.List{}
	list1.First, list1.Last, list1.Current = dt.scanner.NextList()
	list2.First, list2.Last, list2.Current = dt.scanner.NextList()
	dynamic.InsertList(list1, list2)
	return dt.checker.CompareList(list2.First, list2.Last, list2.Current)
}

func (dt *dynamicTask) dynamic69() bool {
	var list1, list2 = &dynamic.List{}, &dynamic.List{}
	list1.First, list1.Last, list1.Current = dt.scanner.NextList()
	list2.First, list2.Last, list2.Current = dt.scanner.NextList()
	dynamic.MoreCurrent(list1, list2)
	return dt.checker.CompareList(list1.First, list1.Last, list1.Current) &&
		dt.checker.CompareList(list2.First, list2.Last, list2.Current)
}

func (dt *dynamicTask) dynamic70() bool {
	var p1, p2, _ = dt.scanner.NextList()
	var barrier = &node.Node{
		Next: p1,
		Prev: p2,
	}
	if p1 != nil {
		p1.Prev = barrier
	} else {
		barrier.Next = barrier
	}
	if p2 != nil {
		p2.Next = barrier
	} else {
		barrier.Prev = barrier
	}
	return dt.checker.CompareBarrierList(barrier, nil)
}

func (dt *dynamicTask) dynamic71() bool {
	var p1, p2 = dt.scanner.NextBarrierList()
	var p3 = &node.Node{}
	p3.Prev, p3.Next = p3, p3
	if p2 != nil && p1 != p2 {
		p3.Prev = p1.Prev
		p3.Next = p2
		p2 = p2.Prev
		p2.Next = p1
		p1.Prev = p2
		p3.Prev.Next = p3
		p3.Next.Prev = p3
	}
	return dt.checker.CompareBarrierList(p1, nil) && dt.checker.CompareBarrierList(p3, nil)
}

func (dt *dynamicTask) dynamic72() bool {
	var p1, _ = dt.scanner.NextBarrierList()
	var p2, _ = dt.scanner.NextBarrierList()
	p1.Prev.Next = p2.Next
	p2.Next.Prev = p1.Prev
	p1.Prev = p2.Prev
	p2.Prev.Next = p1
	return dt.checker.CompareList(p1.Next, p1.Prev, nil)
}

func (dt *dynamicTask) dynamic73() bool {
	var p1, _ = dt.scanner.NextBarrierList()
	var p2, _ = dt.scanner.NextBarrierList()
	p2.Next.Prev = p1.Prev
	p1.Prev.Next = p2.Next
	p2.Next = p1.Next
	p1.Next.Prev = p2
	return dt.checker.CompareList(p2.Next, p2.Prev, nil)
}

func (dt *dynamicTask) dynamic74() bool {
	var listB = dynamic.BarrierList{}
	listB.Barrier, listB.Current = dt.scanner.NextBarrierList()
	var n = dt.scanner.NextInt()
	var data int
	for index := 0; index < n; index++ {
		data = dt.scanner.NextInt()
		listB.InsertLast(data)
	}
	return dt.checker.CompareBarrierList(listB.Barrier, listB.Current)
}

func (dt *dynamicTask) dynamic75() bool {
	var listB = dynamic.BarrierList{}
	listB.Barrier, listB.Current = dt.scanner.NextBarrierList()
	var n = dt.scanner.NextInt()
	var data int
	for index := 0; index < n; index++ {
		data = dt.scanner.NextInt()
		listB.InsertFirst(data)
	}
	return dt.checker.CompareBarrierList(listB.Barrier, listB.Current)
}

func (dt *dynamicTask) dynamic76() bool {
	var listB = dynamic.BarrierList{}
	listB.Barrier, listB.Current = dt.scanner.NextBarrierList()
	var data int
	for index := 0; index < 5; index++ {
		data = dt.scanner.NextInt()
		listB.InsertBefore(data)
	}
	return dt.checker.CompareBarrierList(listB.Barrier, listB.Current)
}

func (dt *dynamicTask) dynamic77() bool {
	var listB = dynamic.BarrierList{}
	listB.Barrier, listB.Current = dt.scanner.NextBarrierList()
	var data int
	for index := 0; index < 5; index++ {
		data = dt.scanner.NextInt()
		listB.InsertAfter(data)
	}
	return dt.checker.CompareBarrierList(listB.Barrier, listB.Current)
}

func (dt *dynamicTask) dynamic78() bool {
	var listB = dynamic.BarrierList{}
	listB.Barrier, listB.Current = dt.scanner.NextBarrierList()
	listB.ToFirst()
	var count = 0
	for !listB.IsBarrier() {
		count++
		if count > 100 {
			break
		}
		if count%2 != 0 {
			listB.SetData(0)
		}
		listB.ToNext()
	}
	return dt.checker.CompareInt(count) && dt.checker.CompareBarrierList(listB.Barrier, listB.Current)
}

func (dt *dynamicTask) dynamic79() bool {
	var listB = dynamic.BarrierList{}
	listB.Barrier, listB.Current = dt.scanner.NextBarrierList()
	listB.ToLast()
	var count = 0
	var data int
	for !listB.IsBarrier() {
		count++
		data = listB.GetData()
		if data%2 == 0 {
			if !dt.checker.CompareInt(data) {
				return false
			}
		}
		listB.ToPrev()
	}
	return dt.checker.CompareInt(count)
}

func (dt *dynamicTask) dynamic80() bool {
	var listB = dynamic.BarrierList{}
	listB.Barrier, listB.Current = dt.scanner.NextBarrierList()
	var data int
	for index := 0; !listB.IsBarrier() && index < 5; index++ {
		data = listB.DeleteCurrent()
		if !dt.checker.CompareInt(data) {
			return false
		}
	}
	return dt.checker.CompareBarrierList(listB.Barrier, listB.Current)
}
