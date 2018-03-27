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
	IsTask   bool
	tree.TreeNode
}

func NewPiece(inWBS *WBS, Name string, Serial int, ID, ParentID string) *Piece {
	p := new(Piece)
	p.Name = Name
	p.Serial = Serial
	p.ID = ID
	inWBS.id2piece[ID] = p
	if ParentID != "" {
		q := inWBS.ID2Piece(ParentID)
		q.AttachChild(p)
	}
	p.ParentID = ParentID
	return p
}

func (p *Piece) String() string {
	return fmt.Sprintf("%s (%s,%d)", p.Name, p.ID, p.Serial)
}

func Pretty(p tree.TreeAble) string {
	return tree.StringChildren(p, "   ")
}

type WBS struct {
	tree.Treeshaped
	id2piece map[string]*Piece
}

func (w *WBS) ID2Piece(id string) *Piece {
	return w.id2piece[id]
}

func NewWBS() {
	w := new(WBS)
	w.id2piece = map[string]*Piece{}
}

func (w *WBS) String() string {
	return Pretty(w.Trunk)
}
