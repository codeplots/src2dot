package main

import (
    "fmt"
    "github.com/codeplots/src2dot/testdata/go/import/sweets"
    f "github.com/codeplots/src2dot/testdata/go/import/fruit"
)

import "path"

func main() {
    fmt.Println(path.Join(".", "import"))
    sweets.Log()
    f.Print()
}
