// +build OMIT

package main

import (
	"fmt"
	"sort"
)

// Box START OMIT
type Box struct {
	width, height int
}

// Box END OMIT

// main START OMIT
type BoxByArea []Box

func (a Box) Area() int                { return a.width * a.height }
func (a BoxByArea) Len() int           { return len(a) }
func (a BoxByArea) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BoxByArea) Less(i, j int) bool { return a[i].Area() < a[j].Area() }

func main() {
	boxes := BoxByArea([]Box{
		Box{12, 8},
		Box{2, 2},
		Box{3, 2},
		Box{4, 2},
		Box{6, 6},
		Box{8, 9},
		Box{10, 8},
	})
	fmt.Println(boxes)
	sort.Sort(boxes)
	fmt.Println(boxes)
}

// main END OMIT
