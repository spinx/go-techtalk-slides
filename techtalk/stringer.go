// +build OMIT

package main

import "fmt"

type Box struct {
	width, height int
}

// Stringer START OMIT
// Stringer is implemented by any value that has a String method,
// which defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print
type Stringer interface {
	String() string
}

// Stringer END OMIT

// main START OMIT
func main() {
	s := Box{12, 8}
	fmt.Println(s)
}

// main END OMIT
