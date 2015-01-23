// +build OMIT

package main

import "fmt"

// Box START OMIT
type Box struct {
	width, height int
}

// Box END OMIT

// String START OMIT
func (b Box) String() string {
	out := ""
	for i := 1; i <= b.height; i++ {
		for j := 1; j <= b.width; j++ {
			if i == 1 || i == b.height || j == 1 || j == b.width {
				out += ". "
			} else {
				out += "  "
			}
		}
		out += "\n"
	}
	return out
}

// String END OMIT

// main START OMIT
func main() {
	b := Box{12, 8}
	fmt.Println(b)
}

// main END OMIT
