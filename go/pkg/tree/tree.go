package tree

import (
	"fmt"
	"math/rand"

	"github.com/shahzodshafizod/ptaskbook/go/pkg/node"
)

type Tree struct {
	Root      *node.Node
	Current   *node.Node
	field     *Field
	level     int
	NodeCount int
}

func (t *Tree) Make() {
	if t.Root == nil {
		t.Root = new(node.Node)
		t.NodeCount++
		fmt.Print("root's Data: ")
		fmt.Scan(&(t.Root.Data))
		t.Current = t.Root
	}
	var answer int
	fmt.Print("Where to go? (0-Exit; 1-Left; 2-Right; 3-Parent):\t")
	fmt.Scan(&answer)
	if answer != 1 && answer != 2 {
		switch answer {
		case 0:
			t.Current = t.Root
			return
		case 3:
			t.Current = t.Current.Parent
		}
		t.Make()
		return
	}
	var newNode *node.Node = new(node.Node)
	t.NodeCount++
	fmt.Print("Data = ")
	fmt.Scan(&(newNode.Data))
	newNode.Parent = t.Current
	if answer == 1 {
		t.Current.Left = newNode
	} else {
		t.Current.Right = newNode
	}
	t.Current = newNode
	t.Make()
}

func (t *Tree) AutoMake(count int) {
	if count <= 0 {
		t.Current = t.Root
		return
	}
	if t.Root == nil {
		t.Root = new(node.Node)
		t.Root.Data = rand.Intn(90) + 10
		t.NodeCount++
		count--
		t.Current = t.Root
		if count <= 0 {
			return
		}
	}
	var direction int
	for {
		direction = rand.Intn(3) + 1
		if !(direction == 3 && t.Current.Parent == nil) {
			break
		}
	}
	var exit = false
	switch direction {
	case 1:
		if t.Current.Left != nil {
			t.Current = t.Current.Left
			exit = true
		}
	case 2:
		if t.Current.Right != nil {
			t.Current = t.Current.Right
			exit = true
		}
	case 3:
		t.Current = t.Current.Parent
		exit = true
	}
	if exit {
		t.AutoMake(count)
		return
	}
	var newNode *node.Node = new(node.Node)
	t.NodeCount++
	count--
	newNode.Data = rand.Intn(90) + 10
	newNode.Parent = t.Current
	if direction == 1 {
		t.Current.Left = newNode
	} else {
		t.Current.Right = newNode
	}
	t.Current = newNode
	t.AutoMake(count)
}

func (t *Tree) GetNodeCount() int {
	return t.NodeCount
}

func (t *Tree) GetLevel() int {
	if t.level == 0 {
		t.setLevel(0)
	}
	return t.level
}

func (t *Tree) setLevel(level int) {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	if level > t.level {
		t.level = level
	}
	var current *node.Node = t.Current
	t.Current = current.Left
	t.setLevel(level + 1)
	t.Current = current.Right
	t.setLevel(level + 1)
}

func (t *Tree) Display() {
	if t.Root == nil {
		return
	}
	if t.field == nil {
		t.field = new(Field)
		t.field.Init(t.Root, t.GetLevel()+1, t.GetNodeCount())
	}
	t.field.Display()
}

func (t *Tree) GetDataCount(data int) int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var count = 0
	if t.Current.Data == data {
		count++
	}
	t.Current = current.Left
	count += t.GetDataCount(data)
	t.Current = current.Right
	count += t.GetDataCount(data)
	return count
}

func (t *Tree) GetDataSum() int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	sum := current.Data
	t.Current = current.Left
	sum += t.GetDataSum()
	t.Current = current.Right
	sum += t.GetDataSum()
	return sum
}

func (t *Tree) GetLeftCount() int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var count = 0
	if current.Parent != nil && current.Parent.Left == current {
		count++
	}
	t.Current = current.Left
	count += t.GetLeftCount()
	t.Current = current.Right
	count += t.GetLeftCount()
	return count
}

func (t *Tree) GetLeafCount() int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var count = 0
	if current.Left == nil && current.Right == nil {
		count++
	}
	t.Current = current.Left
	count += t.GetLeafCount()
	t.Current = current.Right
	count += t.GetLeafCount()
	return count
}

func (t *Tree) GetLeafSum() int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var sum = 0
	if current.Left == nil && current.Right == nil {
		sum += current.Data
	}
	t.Current = current.Left
	sum += t.GetLeafSum()
	t.Current = current.Right
	sum += t.GetLeafSum()
	return sum
}

func (t *Tree) GetRightLeafCount() int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var count int = 0
	if current.Left == nil && current.Right == nil &&
		current.Parent != nil && current.Parent.Right == current {
		count++
	}
	t.Current = current.Left
	count += t.GetRightLeafCount()
	t.Current = current.Right
	count += t.GetRightLeafCount()
	return count
}

func (t *Tree) LevelToArray(array []int, level int, toCount bool) {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	var current *node.Node = t.Current
	if toCount {
		array[level]++
	} else {
		array[level] += current.Data
	}
	t.Current = current.Left
	t.LevelToArray(array, level+1, toCount)
	t.Current = current.Right
	t.LevelToArray(array, level+1, toCount)
}

func (t *Tree) Infix() {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	var current *node.Node = t.Current
	t.Current = current.Left
	t.Infix()
	fmt.Printf("%d\t", current.Data)
	t.Current = current.Right
	t.Infix()
}

func (t *Tree) Prefix() {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	var current *node.Node = t.Current
	fmt.Printf("%d\t", current.Data)
	t.Current = current.Left
	t.Prefix()
	t.Current = current.Right
	t.Prefix()
}

func (t *Tree) Postfix() {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	var current *node.Node = t.Current
	t.Current = current.Left
	t.Postfix()
	t.Current = current.Right
	t.Postfix()
	fmt.Printf("%d\t", current.Data)
}

func (t *Tree) InfixToN(index *int, n int) {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	var current *node.Node = t.Current
	t.Current = current.Left
	t.InfixToN(index, n)
	if *index < n {
		fmt.Printf("%d\t", current.Data)
		(*index)++
	} else {
		t.Current = t.Root
		return
	}
	t.Current = current.Right
	t.InfixToN(index, n)
}

func (t *Tree) PostfixFromN(index *int, n int) {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	var current *node.Node = t.Current
	t.Current = current.Left
	t.PostfixFromN(index, n)
	t.Current = current.Right
	t.PostfixFromN(index, n)
	(*index)++
	if *index >= n {
		fmt.Printf("%d\t", current.Data)
	}
}

func (t *Tree) PrefixBetween(index *int, n1 int, n2 int) {
	if t.Current == nil {
		t.Current = t.Root
		return
	}
	(*index)++
	if *index >= n1 && *index <= n2 {
		fmt.Printf("%d\t", t.Current.Data)
	} else if *index > n2 {
		t.Current = t.Root
		return
	}
	var current *node.Node = t.Current
	t.Current = current.Left
	t.PrefixBetween(index, n1, n2)
	t.Current = current.Right
	t.PrefixBetween(index, n1, n2)
}

func (t *Tree) PrintLevel(index int, level int) int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var count = 0
	if index == level {
		count++
		fmt.Printf("%d\t", t.Current.Data)
	}
	var current *node.Node = t.Current
	t.Current = current.Left
	count += t.PrintLevel(index+1, level)
	t.Current = current.Right
	count += t.PrintLevel(index+1, level)
	return count
}

func (t *Tree) GetMaxData() int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var maximal int = current.Data
	if current.Left != nil {
		t.Current = current.Left
		data := t.GetMaxData()
		if data > maximal {
			maximal = data
		}
	}
	if current.Right != nil {
		t.Current = current.Right
		data := t.GetMaxData()
		if data > maximal {
			maximal = data
		}
	}
	t.Current = t.Root
	return maximal
}

func (t *Tree) GetMinData() int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var minimal int = current.Data
	if current.Left != nil {
		t.Current = current.Left
		data := t.GetMinData()
		if data < minimal {
			minimal = data
		}
	}
	if current.Right != nil {
		t.Current = current.Right
		data := t.GetMinData()
		if data < minimal {
			minimal = data
		}
	}
	t.Current = t.Root
	return minimal
}

func (t *Tree) GetLeafDataCount(data int) int {
	if t.Current == nil {
		t.Current = t.Root
		return 0
	}
	var current *node.Node = t.Current
	var count = 0
	if current.Data == data && current.Left == nil && current.Right == nil {
		count++
	}
	t.Current = current.Left
	count += t.GetLeafDataCount(data)
	t.Current = current.Right
	count += t.GetLeafDataCount(data)
	return count
}

func GetCount(node *node.Node, data int, direction int, options ...string) int {
	if node == nil {
		return 0
	}
	var count = 0
	for _, option := range options {
		switch option {
		case "data":
			if node.Data == data {
				count++
			}
		case "left":
			if direction == 1 {
				count++
			}
		case "right":
			if direction == 2 {
				count++
			}
		case "leaf":
			if node.Left == nil && node.Right == nil {
				count++
			}
		case "all":
			count++
		}
	}
	if count == len(options) {
		count = 1
	} else {
		count = 0
	}
	return count + GetCount(node.Left, data, 1, options...) + GetCount(node.Right, data, 2, options...)
}

func GetSum(node *node.Node, data int, direction int, options ...string) int {
	if node == nil {
		return 0
	}
	var sum = 0
	for _, option := range options {
		switch option {
		case "data":
			if node.Data == data {
				sum += node.Data
			}
		case "left":
			if direction == 1 {
				sum += node.Data
			}
		case "right":
			if direction == 2 {
				sum += node.Data
			}
		case "leaf":
			if node.Left == nil && node.Right == nil {
				sum += node.Data
			}
		case "all":
			sum += node.Data
		}
	}
	if sum == len(options)*node.Data {
		sum = node.Data
	} else {
		sum = 0
	}
	return sum + GetSum(node.Left, data, 1, options...) + GetSum(node.Right, data, 2, options...)
}
