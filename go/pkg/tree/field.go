package tree

import (
	"fmt"

	"github.com/shahzodshafizod/ptaskbook/go/pkg/node"
)

const (
	SPACE  int8 = 0
	LEFT   int8 = 1
	CENTER int8 = 2
	RIGHT  int8 = 3
	DATA   int8 = 4
)

type Field struct {
	nodes     [][]*node.Node
	chars     [][]int8
	height    int
	width     int
	cellWidth int
	left      string
	center    string
	right     string
}

func (f *Field) Init(n *node.Node, height int, width int) {
	f.height = height
	f.width = width
	f.left = string(rune(9484))
	f.center = string(rune(9472))
	f.right = string(rune(9488))
	f.cellWidth = 4
	f.nodes = make([][]*node.Node, f.height)
	f.chars = make([][]int8, f.height)
	for row := 0; row < f.height; row++ {
		f.nodes[row] = make([]*node.Node, f.width)
		f.chars[row] = make([]int8, f.width)
	}
	var column int = 0
	f.fillNodes(n, 0, &column)
	f.fillChars()
}

func (f *Field) fillNodes(node *node.Node, row int, col *int) {
	if node == nil {
		return
	}
	f.fillNodes(node.Left, row+1, col)
	f.nodes[row][*col] = node
	(*col)++
	f.fillNodes(node.Right, row+1, col)
}

func (f *Field) fillChars() {
	for row := 0; row < f.height; row++ {
		for col := 0; col < f.width; col++ {
			if f.nodes[row][col] == nil {
				continue
			}
			f.chars[row][col] = DATA
			if f.nodes[row][col].Left != nil {
				for k := col - 1; k >= 0; k-- {
					if f.nodes[row+1][k] != nil {
						f.chars[row][k] = LEFT
						break
					}
					f.chars[row][k] = CENTER
				}
			}
			if f.nodes[row][col].Right != nil {
				for col++; col < f.width; col++ {
					if f.nodes[row+1][col] != nil {
						f.chars[row][col] = RIGHT
						break
					}
					f.chars[row][col] = CENTER
				}
			}
		}
	}
}

func (f Field) getNextNode(row int, col int) *node.Node {
	for col++; col < f.width; col++ {
		if f.nodes[row][col] != nil {
			return f.nodes[row][col]
		}
	}
	return nil
}

func getLength(number int) int {
	var len int = 0
	if number < 0 {
		len++
		number *= -1
	}
	for number > 0 {
		len++
		number /= 10
	}
	return len
}

func printChars(str string, count int) {
	for count > 0 {
		fmt.Print(str)
		count--
	}
}

func (f Field) Display() {
	var (
		nextNode           *node.Node
		data, spaces, half int
	)
	for row := range f.chars {
		for col := range f.chars[row] {
			switch f.chars[row][col] {
			case SPACE:
				printChars(" ", f.cellWidth)
			case LEFT:
				nextNode = f.getNextNode(row, col)
				data = nextNode.Data
				spaces = f.cellWidth - getLength(data)
				half = spaces / 2
				printChars(" ", half)
				printChars(f.left, 1)
				printChars(f.center, f.cellWidth-1)
			case CENTER:
				printChars(f.center, f.cellWidth)
			case RIGHT:
				printChars(f.center, f.cellWidth-1)
				printChars(f.right, 1)
				printChars(" ", spaces-half)
			case DATA:
				if f.nodes[row][col].Left == nil {
					data = f.nodes[row][col].Data
					spaces = f.cellWidth - getLength(data)
					half = spaces / 2
					printChars(" ", half)
				}
				fmt.Print(data)
				if f.nodes[row][col].Right == nil {
					printChars(" ", spaces-half)
				}
			}
		}
		fmt.Println()
	}
}
