package main

import (
    "github.com/codeplots/src2dot/testdata/go/composition/plants"
    "fmt"
)

func main() {
    fmt.Println("main")
    someTree := plants.Tree{}
    someTree.Grow()
}

