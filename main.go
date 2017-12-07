package main

import (
	"fmt"
	db "github.com/codeplots/src2dot/database"
)

func main() {
	opts := parseOptions()
	fmt.Println("Options: ", opts)
	// cscope -kbc
	// ctags --all-kinds=* --fields=* -R --extras=+frF .
	myDb := db.FromCtags(opts.ctags)
	myDb.AddCscope(opts.cscope)
	fmt.Println(myDb)
}
