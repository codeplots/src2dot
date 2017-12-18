package plants

import "fmt"

type Plant struct {
    age int
}

type Tree struct {
    Plant
    height int
}

func (t *Tree) Grow() {
    fmt.Println("Tree grows")
}
