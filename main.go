package main

import (
	"fmt"

	ivan "github.com/robertely/fisher/checkers/crazyivan"
)

// var x Checker

func main() {
	x := ivan.New()
	x.Name = "Ivan_1"
	fmt.Println(x.Check()) // No work

	fmt.Println(x)
}
