package main

import (
	"fmt"

	"github.com/aprice2704/seldon/wbs"
)

func main() {

	test := wbs.Piece{Name: "Aldrin Base", Serial: 1000}
	fmt.Println(test)
	test2 := wbs.Piece{Name: "Eagle Landing Pad", Serial: 1100}
	fmt.Println(test2)

	fmt.Println("---------")

	test.AttachChild(test2)
	fmt.Println(test)
}
