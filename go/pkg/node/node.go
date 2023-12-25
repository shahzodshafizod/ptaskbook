package node

type Node struct {
	Data   int
	Prev   *Node
	Next   *Node
	Parent *Node
	Left   *Node
	Right  *Node
}
