package main

import (
	"fmt"

	"github.com/aprice2704/seldon/wbs"
)

func main() {

	byid := map[string]*wbs.Piece{}

	pieces := map[string]wbs.Piece{ // Just a temp structure to read in values
		"A001": {Name: "Aldrin Base", Serial: 1000, ID: "A001"},
		"A002": {Name: "Eagle Landing Pad", Serial: 1100, ID: "A002", ParentID: "A001"},
		"A003": {Name: "Foundation", Serial: 1200, ID: "A003", ParentID: "A002"},
		"A004": {Name: "Perimeter berm", Serial: 1300, ID: "A003", ParentID: "A002"},
		"A005": {Name: "Main Building", Serial: 1050, ID: "A004", ParentID: "A001"},
	}

	for k, v := range pieces {
		byid[k] = wbs.NewPiece(v.Name, v.Serial, v.ID, v.ParentID)
		fmt.Println(byid[k])
	}

	fmt.Println("----")
	for _, v := range byid {
		p, ok := byid[v.ParentID]
		if ok {
			fmt.Println("Attach: ", v.String(), " to ", p.String())
			p.AttachChild(v)
		}
	}

	fmt.Println("----")
	fmt.Println(wbs.Pretty(byid["A001"]))

}
