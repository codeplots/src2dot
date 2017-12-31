package main

import (
	"flag"
	"fmt"
        "encoding/json"
	db "github.com/codeplots/src2dot/database"
	"github.com/codeplots/src2dot/graph"
	"io/ioutil"
	"os"
)

const (
	DEP_FILENAME   string      = "dependency_graph"
	CLASS_FILENAME string      = "class_diagram"
	CALL_FILENAME  string      = "call_graph"
	PERMISSION     os.FileMode = 0644
)

func main() {
	opts := parseOptions()

	store := db.FromCtags(opts.ctags)
	store.AddCscope(opts.cscope)

	if !(opts.dep || opts.class || opts.call) {
            for _, filename := range []string{
                    DEP_FILENAME,
                    CLASS_FILENAME,
                    CALL_FILENAME,
                } {
                var g graph.Graph

                switch filename {
                case DEP_FILENAME:
                    g, _ = graph.DependencyGraph(store)
                case CLASS_FILENAME:
                    g, _ = graph.ClassDiagram(store)
                case CALL_FILENAME:
                }

                var b []byte
                var ext string
                switch opts.raw {
                case true:
                    b, _ = json.Marshal(g)
                    ext = ".json"
                case false:
                    b = []byte(g.ToDot())
                    ext = ".dot"
                }
		err := ioutil.WriteFile(filename + ext, b, PERMISSION)
		if err != nil {
			panic(err)
		}
            }
            return
	}

        graphs := []graph.Graph{}

	if opts.dep {
		dep, _ := graph.DependencyGraph(store)
                graphs = append(graphs, dep)
	}
	if opts.class {
		class, _ := graph.ClassDiagram(store)
                graphs = append(graphs, class)
	}
        switch opts.raw {
        case true:
            bytes, _ := json.Marshal(graphs)
            fmt.Println(string(bytes))
        case false:
            for _, g := range graphs {
                fmt.Println(g.ToDot())
            }
        }

}

type Options struct {
	ctags  string
	cscope string
	raw    bool
	dep    bool
	class  bool
	call   bool
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
