package wbs

import (
	"fmt"

	"github.com/aprice2704/seldon/tree"
)

type Piece struct {
	Name   string
	Serial int
	ID     string
	tree.TreeNode
}

func (p Piece) String() string {
	s := fmt.Sprintf("%s (%d)", p.Name, p.Serial)
	chilluns := p.Children()
	fmt.Println("chilluns:", chilluns)
	for _, c := range chilluns {
		fmt.Println(c.(Piece))
		s += "\n   " + c.(Piece).String()
	}
	return s
}
