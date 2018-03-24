package wbs

import (
	"fmt"

	"github.com/aprice2704/seldon/tree"
)

type Piece struct {
	Name     string
	Serial   int
	ID       string
	ParentID string
	tree.TreeNode
}

func (p *Piece) String() string {
	return fmt.Sprintf("%s (%s,%d)", p.Name, p.ID, p.Serial)
}

func Pretty(p tree.Treeish) string {
	return tree.StringChildren(p, "   ")
}
