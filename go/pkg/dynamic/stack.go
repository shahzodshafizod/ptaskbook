package dynamic

import "github.com/shahzodshafizod/ptaskbook/go/pkg/node"

type Stack struct {
	Top *node.Node
}

func (s *Stack) Push(d int) {
	s.Top = &node.Node{
		Data: d,
		Next: s.Top,
	}
}

func (s *Stack) Pop() int {
	if s.Top == nil {
		return 0
	}
	var data = s.Top.Data
	s.Top = s.Top.Next
	return data
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}
