package calendar

import (
	"fmt"
	"testing"
	"time"
)

func TestDayLayer(t *testing.T) {
	const shortForm = "2006-Jan-02"
	dls, _ := time.Parse(shortForm, "2018-Mar-01")
	dle, _ := time.Parse(shortForm, "2018-Mar-20")
	dl := NewCalLayer(dls, dle)
	fmt.Println(dl.String())
	// Output:
	// Layer, Starts: 62 Ends: 66
	aus, _ := time.Parse(shortForm, "2018-Mar-02")
	aue, _ := time.Parse(shortForm, "2018-Mar-19")
	if !dl.IsAllUsable(aus, aue) {
		t.Error("Useable time test 1")
	}
}
