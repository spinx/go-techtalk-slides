// +build OMIT

package main

import "fmt"

// Box START OMIT
type Box struct {
	width, height int
}

// Box END OMIT

// String START OMIT
func (b Box) Describe() string {
	return fmt.Sprintf("This is a box of %d by %d", b.width, b.height)
}

// String END OMIT

// main START OMIT
func main() {
	b := Box{8, 8}
	fmt.Println(b.Describe())
}

// main END OMIT
