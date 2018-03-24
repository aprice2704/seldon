package tree

import "fmt"

type Treeish interface {
	AttachChild(c Treeish)
	DetachChild(c Treeish) error
	Children() []Treeish
	Parent() Treeish
	String() string
}

type TreeNode struct {
	parent   Treeish
	children []Treeish
}

func StringChildren(n Treeish, indent string) string {
	s := n.String()
	chilluns := n.Children()
	for _, c := range chilluns {
		s += "\n" + indent + StringChildren(c, indent+"   ")
	}
	return s
}

func (node *TreeNode) AttachChild(c Treeish) {
	if node.children == nil {
		node.children = make([]Treeish, 0)
	}
	node.children = append(node.children, c)
}

func (node *TreeNode) DetachChild(c Treeish) error {
	fmt.Println("Detaching")
	return nil
}

func (node *TreeNode) Children() []Treeish {
	return node.children
}

func (node *TreeNode) Parent() Treeish {
	return node.parent
}

func (t TreeNode) String() string {
	return "A plain treenode"
}
