package tree

type Type int

const (
	Concat Type = 1 + iota
	Or
	Star
	Sharp
	Group
	Reference
	Repeat
	LeafNode
	Bracket
)

type Node struct {
	Id       int // Leaf type >= 1, another = -1
	Parent   *Node
	Left     *Node
	Right    *Node
	Type     Type
	Val      string
	FirstPos []int
	LastPos  []int
	Nullable bool
}

type NodePrint struct {
	ChildState     *NodePrint
	PrintLastChild bool
}

type Ast struct {
	Root      *Node
	FollowPos [][]int
	ID map[int]string
	LeafNodes map[string][]int
}

var rootState *NodePrint = nil