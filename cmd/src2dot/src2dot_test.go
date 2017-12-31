package main

import (
	"fmt"
	db "github.com/codeplots/src2dot/database"
	"github.com/codeplots/src2dot/graph"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
	"testing"
)

const (
	DEP_GRAPH_FILENAME  string = "dependency_graph.dot"
	CALL_GRAPH_FILENAME string = "call_graph.dot"
	CLASS_DIAG_FILENAME string = "class_diagram.dot"
)

func TestDotting(t *testing.T) {
	dirs := getTestDataDirs("../../testdata")
	for _, dir := range dirs {
		testDb := db.FromCtags(path.Join(dir, "tags"))
		testDb.AddCscope(path.Join(dir, "cscope.out"))
		files, _ := ioutil.ReadDir(dir)
		for _, f := range files {
			if !strings.HasSuffix(f.Name(), ".dot") {
				continue
			}
			t.Run(dir, func(t *testing.T) {
				var g graph.Graph
				var expected []byte

				switch n := f.Name(); {
				case n == DEP_GRAPH_FILENAME:
					g, _ = graph.DependencyGraph(testDb)
					expected, _ = ioutil.ReadFile(path.Join(
						dir, DEP_GRAPH_FILENAME))
				case n == CALL_GRAPH_FILENAME:
					fmt.Println("No support for call graphs yet")
				case n == CLASS_DIAG_FILENAME:
					g, _ = graph.ClassDiagram(testDb)
					expected, _ = ioutil.ReadFile(path.Join(
						dir, CLASS_DIAG_FILENAME))
				case strings.HasSuffix(n, ".dot"):
					fmt.Printf("Warning: Could not recognize %s\n", n)
					t.Skip()
				}
				actual := g.ToDot()
				isEqual := hasSameEdges(string(expected), actual) && hasSameNodes(string(expected), actual)
				if !isEqual {
					t.Error("The output graph (1) does not match the expected graph (2):\n", "(1)", actual, "\n(2)", string(expected))
				}
			})
		}
	}
}

func hasSameNodes(dot1 string, dot2 string) bool {
	if !aContainsAllNodesFromB(dot1, dot2) {
		return false
	}
	if !aContainsAllNodesFromB(dot2, dot1) {
		return false
	}
	return true
}

func aContainsAllNodesFromB(a string, b string) bool {
	bLines := strings.Split(b, "\n")
	r := regexp.MustCompile(`^\s*"(.*)"\s*\[\s*label=\"(.*)"\s*\]\s*$`)
	for _, line := range bLines {
		matches := r.FindStringSubmatch(line)
		if len(matches) == 3 {
			id := matches[1]
			label := matches[2]
			if !containsNode(a, id, label) {
				return false
			}
		}
	}
	return true
}

func containsNode(dot string, id string, label string) bool {
	r := regexp.MustCompile(`\s*"` +
		regexp.QuoteMeta(id) + `"\s*\[\s*label="` +
		regexp.QuoteMeta(label) + `"\s*\]\s*\n`)
	found := r.FindStringIndex(dot)
	if found == nil {
		return false
	}
	return true
}

func hasSameEdges(dot1 string, dot2 string) bool {
	styles := []graph.EdgeStyle{
		graph.LINE,
		graph.ARROW,
		graph.HOLLOW_DIAMOND,
		graph.FILLED_DIAMOND,
	}
	for _, s := range styles {
		if !aContainsAllEdgesFromB(dot1, dot2, s) {
			return false
		}
		if !aContainsAllEdgesFromB(dot2, dot1, s) {
			return false
		}
	}
	return true
}

func aContainsAllEdgesFromB(a string, b string, style graph.EdgeStyle) bool {
	bLines := strings.Split(b, "\n")
	var r *regexp.Regexp
	switch style {
	case graph.LINE:
		r = regexp.MustCompile(`^\s*"(.*)"\s*->\s*"(.*)"\s*\[dir="?none"?\]\s*$`)
	case graph.HOLLOW_DIAMOND:
		r = regexp.MustCompile(`^\s*"(.*)"\s*->\s*"(.*)"\s*\[arrowhead="?ediamond"?\]\s*$`)
	case graph.FILLED_DIAMOND:
		r = regexp.MustCompile(`^\s*"(.*)"\s*->\s*"(.*)"\s*\[arrowhead="?diamond"?\]\s*$`)
	default:
		r = regexp.MustCompile(`^\s*"(.*)"\s*->\s*"(.*)"\s*$`)
	}
	for _, line := range bLines {
		matches := r.FindStringSubmatch(line)
		if len(matches) == 3 {
			source := matches[1]
			target := matches[2]
			if !containsEdge(a, source, target, style) {
				return false
			}
		}
	}
	return true
}

func containsEdge(dot string, source string, target string, style graph.EdgeStyle) bool {
	var r *regexp.Regexp
	switch style {
	case graph.LINE:
		r = regexp.MustCompile(`\s*"` +
			regexp.QuoteMeta(source) + `"\s*->\s*"` +
			regexp.QuoteMeta(target) + `"\s*\[dir="?none"?\]\s*\n`)
	case graph.HOLLOW_DIAMOND:
		r = regexp.MustCompile(`\s*"` +
			regexp.QuoteMeta(source) + `"\s*->\s*"` +
			regexp.QuoteMeta(target) + `"\s*\[arrowhead="?ediamond"?\]\s*\n`)
	case graph.FILLED_DIAMOND:
		r = regexp.MustCompile(`\s*"` +
			regexp.QuoteMeta(source) + `"\s*->\s*"` +
			regexp.QuoteMeta(target) + `"\s*\[arrowhead="?diamond"?\]\s*\n`)
	default:
		r = regexp.MustCompile(`\s*"` +
			source + `"\s*->\s*"` +
			target + `"\s*\n`)
	}
	found := r.FindStringIndex(dot)
	if found == nil {
		return false
	}
	return true
}

func getTestDataDirs(baseDir string) []string {
	dirs := []string{}

	langFolders, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatalf("Could not open '%s'", baseDir)
	}

	for _, lang := range langFolders {
		if !lang.IsDir() {
			continue
		}
		langPath := path.Join(path.Join(baseDir, lang.Name()))
		codeFolders, err := ioutil.ReadDir(langPath)
		if err != nil {
			log.Fatalf("Could not open '%s'", langPath)
		}

		for _, code := range codeFolders {
			if !code.IsDir() {
				continue
			}
			dirs = append(dirs, path.Join(langPath, code.Name()))
		}
	}
	return dirs
}
