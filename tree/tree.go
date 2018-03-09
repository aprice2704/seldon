package tree

import "fmt"

type Treeish interface {
	AttachChild(c Treeish)
	DetachChild(c Treeish) error
	Children() []Treeish
	Parent() Treeish
}

type TreeNode struct {
	parent   Treeish
	children []Treeish
}

func (node TreeNode) AttachChild(c Treeish) {
	fmt.Println("Attaching")
	if node.children == nil {
		fmt.Println("make children")
		node.children = make([]Treeish, 0)
	}
	fmt.Println(node)
	node.children = append(node.children, c)
	fmt.Println(node)
}

func (node TreeNode) DetachChild(c Treeish) error {
	fmt.Println("Detaching")
	return nil
}

func (node TreeNode) Children() []Treeish {
	return node.children
}

func (node TreeNode) Parent() Treeish {
	return node.parent
}
