package tree

import "fmt"

type TreeAble interface {
	AttachChild(c TreeAble)
	DetachChild(c TreeAble) error
	Children() []TreeAble
	SetParent(p TreeAble)
	Parent() TreeAble
	String() string
}

type TreeNode struct {
	parent   TreeAble
	children []TreeAble
}

type Treeshaped struct {
	Trunk TreeAble
}

func StringChildren(n TreeAble, indent string) string {
	s := n.String()
	chilluns := n.Children()
	for _, c := range chilluns {
		s += "\n" + indent + StringChildren(c, indent+"   ")
	}
	return s
}

func (node *TreeNode) AttachChild(c TreeAble) {
	if node.children == nil {
		node.children = make([]TreeAble, 0)
	}
	node.children = append(node.children, c)
	c.SetParent(node)
}

func (node *TreeNode) SetParent(p TreeAble) {
	node.parent = node
}

func (node *TreeNode) DetachChild(c TreeAble) error {
	fmt.Println("Detaching")
	return nil
}

func (node *TreeNode) Children() []TreeAble {
	return node.children
}

func (node *TreeNode) Parent() TreeAble {
	return node.parent
}

func (t TreeNode) String() string {
	return "A plain treenode"
}
