package fruit

import (
    "fmt"
    "strings"
    "github.com/codeplots/src2dot/testdata/go/import/sweets"
)

func Print() {
    fmt.Println(strings.TrimPrefix("Apple", ""))
    sweets.Log()
}
