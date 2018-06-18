package main

import (
	"testing"

	"github.com/aprice2704/seldon/wbs"
)

func TestWBS(t *testing.T) {
	mywbs := wbs.NewWBS()
	p := wbs.NewPiece(mywbs, "Testsomething", 42, "A0042", "")
	s := p.String()
	shouldb := "Testsomething (A0042,42)"
	if s != shouldb {
		t.Errorf("WBS: basic object creation failed -- is %s, should be %s", s, shouldb)
	}
}
