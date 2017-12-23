package main

import (
        "flag"
	"fmt"
        "io/ioutil"
        "os"
	db "github.com/codeplots/src2dot/database"
        "github.com/codeplots/src2dot/graph"
)

const (
    DEP_FILENAME        string = "dependency_graph.dot"
    CLASS_FILENAME      string = "class_diagram.dot"
    CALL_FILENAME       string = "call_graph.dot"
    PERMISSION          os.FileMode    = 0644
)

func main() {
	opts := parseOptions()

	store := db.FromCtags(opts.ctags)
	store.AddCscope(opts.cscope)

        if !(opts.dep || opts.class || opts.call) {
            dep, _ := graph.DependencyGraph(store)
            err := ioutil.WriteFile(DEP_FILENAME, []byte(dep.ToDot()), PERMISSION)
            if err != nil {
                panic(err)
            }
            class, _ := graph.ClassDiagram(store)
            err = ioutil.WriteFile(CLASS_FILENAME, []byte(class.ToDot()), PERMISSION)
            if err != nil {
                panic(err)
            }

            return
        }

        if opts.dep {
            dep, _ := graph.DependencyGraph(store)
            fmt.Println(dep.ToDot())
        }
        if opts.class {
            class, _ := graph.ClassDiagram(store)
            fmt.Println(class.ToDot())
        }

}

type Options struct {
	ctags  string
	cscope string
	raw    bool
	dep      bool
	class      bool
	call      bool
}

func parseOptions() Options {
	opts := Options{}
	flag.StringVar(&opts.ctags, "ctags", "./tags",
		"path to ctags output file")
	flag.StringVar(&opts.cscope, "cscope", "./cscope.out",
		"path to cscope output file")
	flag.BoolVar(&opts.dep, "dep", false,
		"print dependency graph to stdout (suppresses all output to files)")
	flag.BoolVar(&opts.class, "class", false,
		"print class diagram to stdout (suppresses all output to files)")
	flag.BoolVar(&opts.call, "call", false,
		"print call graph to stdout (suppresses all output to files)")
	flag.BoolVar(&opts.raw, "raw", false,
		"output raw json graph")
	flag.Parse()
	return opts
}
