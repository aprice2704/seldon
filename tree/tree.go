package tree

import "fmt"

// Simple int based tree structure, designed to apply to any slice of objects -- objects[i]
// The tree only stores and manages the indexes of the real objects within an array/slice
// The data in the tree is arranged as the objects are, thus: Node[i] carries the structrual links
// for objects[i]

// Node tracks a node's relatives
type Node struct {
	Parent   int
	IsChild  bool
	Children []int
}

// Nodes is
type Nodes []*Node

// Tree holds the structure of the whole tree
type Tree struct {
	Root  int
	Nodes Nodes
}

// NewTree creates an empty tree for use with nitems (initially)
func NewTree(nitems int, rootnode int) *Tree {
	return &Tree{Root: rootnode, Nodes: make(Nodes, nitems)}
}

// Attach a child (leaf) to a parent within the tree
func (t Tree) Attach(leaf, branch int) {
	if t.Nodes[branch] == nil {
		t.Nodes[branch] = &Node{Children: []int{leaf}, IsChild: false}
	} else {
		t.Nodes[branch].Children = append(t.Nodes[branch].Children, leaf)
	}
	if t.Nodes[leaf] == nil {
		t.Nodes[leaf] = &Node{Parent: branch, IsChild: true}
	}
	//fmt.Printf("%s\n", t.Nodes[leaf])
}

// NChildren returns the number of children the node has
func (t Tree) NChildren(n int) int {
	if t.Nodes[n] == nil {
		return 0
	}
	return len(t.Nodes[n].Children)
}

// Children returns the children of a node
func (t Tree) Children(ofnode int) []int {
	if t.Nodes[ofnode] == nil {
		return []int{}
	}
	return t.Nodes[ofnode].Children
}

func (t Tree) String() string {
	s := fmt.Sprintf("Tree: root node is #%d, of %d nodes", t.Root, len(t.Nodes))
	s = s + t.Nodes[t.Root].StringR(t.Nodes)
	return s
}

// StringR does a recursive string render
func (n Node) StringR(nodes Nodes) string {
	var s string
	if n.IsChild {
		s = fmt.Sprintf("Node's parent is %d has %d children: %v. ", n.Parent, len(n.Children), n.Children)
	} else {
		s = fmt.Sprintf("Node is root, has %d children: %v. ", len(n.Children), n.Children)
	}
	for i, v := range n.Children {
		s = s + fmt.Sprintf("Child #%d is %d: ", i, v) + nodes[v].StringR(nodes)
	}
	return s
}

func (n Node) String() string {
	var s string
	if n.IsChild {
		s = fmt.Sprintf("Node's parent is %d has %d children: %v. ", n.Parent, len(n.Children), n.Children)
	} else {
		s = fmt.Sprintf("Node is root, has %d children: %v. ", len(n.Children), n.Children)
	}
	return s
}
