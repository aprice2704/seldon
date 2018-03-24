package main

import (
	"fmt"

	"github.com/aprice2704/seldon/wbs"
)

func main() {

	// test2 := &wbs.Piece{Name: "Eagle Landing Pad", Serial: 1100}
	// test3 := &wbs.Piece{Name: "Foundation", Serial: 1200}
	// test4 := &wbs.Piece{Name: "Perimeter berm", Serial: 1300}
	// test5 := &wbs.Piece{Name: "Main Building", Serial: 1050}

	byid := map[string]wbs.Piece{
		"A001": {Name: "Aldrin Base", Serial: 1000, ID: "A001"},
		"A002": {Name: "Eagle Landing Pad", Serial: 1100, ID: "A002", ParentID: "A001"},
		"A003": {Name: "Foundation", Serial: 1200, ID: "A003", ParentID: "A002"},
		"A004": {Name: "Perimeter berm", Serial: 1300, ID: "A003", ParentID: "A002"},
		"A005": {Name: "Main Building", Serial: 1050, ID: "A004", ParentID: "A001"},
	}

	//aldrin := &wbsbits["A001"]

	//fmt.Println(wbsbits)

	//byid := make(map[string]*wbs.Piece) // ID --> address
	//byid[aldrin.ID] = aldrin
	// for _, v := range wbsbits {
	// 	cp := deepcopy.Copy(v).(wbs.Piece)
	// 	byid[v.ID] = cp
	// 	fmt.Println(v.String())
	// }

	fmt.Println("----")
	fmt.Println(byid)

	fmt.Println("----")
	for _, v := range byid {
		p, ok := byid[v.ParentID]
		if ok {
			fmt.Println("Attach: ", v.String(), " to ", p.String())
			//	parents[k] = p
			p.AttachChild(&v)
		}
	}

	// test.AttachChild(test2)
	// test2.AttachChild(test3)
	// test2.AttachChild(test4)
	// test.AttachChild(test5)

	fmt.Println(wbs.Pretty(byid["A001"]))

}
