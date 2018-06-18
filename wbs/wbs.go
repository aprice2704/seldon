package wbs

import (
	"fmt"

	"github.com/aprice2704/seldon/tree"
)

// WBS is a Work Breakdown Structure
type WBS struct {
	Tree     *tree.Tree
	ID       string //
	RootID   string
	ID2Piece map[string]int
	Pieces   Pieces
}

// Piece is a piece of the Work
type Piece struct {
	Name     string
	Serial   int
	ID       string
	ParentID string
	IsTask   bool
}

// Pieces is just a list of piece
type Pieces []Piece

func (p *Piece) String() string {
	return fmt.Sprintf("%s (id:%s, ser:%d) ^%s", p.Name, p.ID, p.Serial, p.ParentID)
}

// NewWBS makes a new one
func NewWBS(rootid string, pieces Pieces) *WBS {
	w := new(WBS)
	w.Pieces = append(w.Pieces, pieces...)
	w.RootID = rootid
	w.ID2Piece = make(map[string]int, 0)
	for i, v := range pieces {
		w.ID2Piece[v.ID] = i
	}
	rootn := w.ID2Piece[rootid]
	w.Tree = tree.NewTree(len(pieces), rootn)
	for i, v := range pieces {
		if len(v.ParentID) > 0 {
			//	fmt.Printf("Attaching %s/%d to %s/%d\n", pieces[i].ID, i, v.ParentID, w.ID2Piece[v.ParentID])
			w.Tree.Attach(i, w.ID2Piece[v.ParentID])
		}
	}
	return w
}

// PieceStringR does a recursive pretty print of a piece and its children
func (w *WBS) PieceStringR(pieceno int, offset, indent string) string {
	//fmt.Printf("WBS Piece #%d\n", pieceno)
	s := offset + w.Pieces[pieceno].String() + "\n"
	//fmt.Println("here1")
	if w.Tree.NChildren(pieceno) > 0 {
		for _, v := range w.Tree.Children(pieceno) {
			//fmt.Printf("Child: %d", v)
			s = s + w.PieceStringR(v, offset+indent, indent)
		}
	}
	return s
}

func (w *WBS) String() string {
	return w.PieceStringR(w.ID2Piece[w.RootID], "", "   ")
}
