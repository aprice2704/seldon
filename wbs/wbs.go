package wbs

import "github.com/aprice2704/seldon/tree"

type Piece struct {
	Name   string
	Serial int
	ID     string
	tree.TreeNode
}
