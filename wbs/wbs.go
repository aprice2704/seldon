package wbs

type Piece struct {
	Name     string
	Serial   int
	ID       string
	Parent   *Piece
	Children []*Piece
}
