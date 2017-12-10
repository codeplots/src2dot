package main

import (
        "flag"
	"fmt"
	db "github.com/codeplots/src2dot/database"
        "github.com/codeplots/src2dot/graph"
)

func main() {
	opts := parseOptions()
	fmt.Println("Options: ", opts)
	// cscope -kbc
	// ctags --all-kinds=* --fields=* -R --extras=+frF .
	myDb := db.FromCtags(opts.ctags)
	myDb.AddCscope(opts.cscope)
        g, _ := graph.DependencyGraph(myDb)
        fmt.Println(g.ToDot())
}

type Options struct {
	ctags  string
	cscope string
	raw    bool
	o      string
}

func parseOptions() Options {
	opts := Options{}
	flag.StringVar(&opts.ctags, "ctags", "./tags",
		"path to tags file")
	flag.StringVar(&opts.cscope, "cscope", "./cscope.out",
		"path to cscope.out file")
	flag.StringVar(&opts.o, "o", "-",
		"output file (- for stdout)")
	flag.BoolVar(&opts.raw, "raw", false,
		"output raw json graph")
	flag.Parse()
	return opts
}
