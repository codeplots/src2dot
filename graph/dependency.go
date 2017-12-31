package graph

import (
	db "github.com/codeplots/src2dot/database"
	"path"
	"regexp"
	"strings"
)

func DependencyGraph(store db.Database) (Graph, error) {
	g := Graph{
		Nodes: []Node{},
		Edges: []Edge{},
		Kind:  DEPENDENCY_GRAPH,
	}

	for _, ref := range store.GetImports() {
		if ref.Module() == "__init__" {
			continue
		}
		id := path.Join(ref.Dir(), ref.Filename())
		g.addNodeIfNotExist(Node{
			Id:       id,
			Label:    id,
			ParentId: ref.Dir(),
		})

		files, _ := store.GetImportedFiles(ref)
		if len(files) == 0 {
			target, isImported := formatSysImport(ref)
			if isImported {
				g.addEdgeIfNotExist(Edge{
					SourceId: id,
					TargetId: target,
				})
			}
			continue
		}
		for _, file := range files {
			fileId := path.Join(file.Dir(), file.Name())
			g.addNodeIfNotExist(Node{
				Id:       fileId,
				Label:    fileId,
				ParentId: file.Dir(),
			})
			g.addEdgeIfNotExist(Edge{
				SourceId: id,
				TargetId: fileId,
			})
		}
	}
	return g, nil
}

func formatSysImport(ref db.ImportRef) (string, bool) {
	format := ""
	switch ref.Language() {
	case db.CPP:
		return ref.Symbol(), true
	case db.GO:
		return ref.Symbol(), true
	case db.PYTHON:
		return formatPythonSysImport(ref)
	}
	return format, false
}

func formatPythonSysImport(ref db.ImportRef) (string, bool) {
	r := regexp.MustCompile(`(?:from\s+(\S+)\s)?\s*import\s+((?:\S+\s*)(?:,\s*\S+)*)(?:\sas.*)?$`)
	matches := r.FindStringSubmatch(ref.LineSrc())
	if len(matches) != 3 {
		return ref.Symbol(), true
	}
	if !strings.Contains(matches[2], ref.Symbol()) {
		return "", false
	}
	res := ref.Symbol()
	if matches[1] != "" {
		res = matches[1] + "." + res
	}
	return res, true
}
