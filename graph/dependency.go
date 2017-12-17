package graph

import (
    "path"
    "strings"
    "regexp"
    db "github.com/codeplots/src2dot/database"
)

func DependencyGraph(store db.Database) (Graph, error) {
    g := Graph {
        nodes: []Node{},
        edges: []Edge{},
    }

    for _, ref := range store.GetImports() {
        if ref.Module() == "__init__" {
            continue
        }
        id := path.Join(ref.Dir(), ref.Filename())
        g.addNodeIfNotExist(Node{
            id: id,
            label: id,
            parentId: ref.Dir(),
        })

        files, _ := store.GetImportedFiles(ref)
        if len(files) == 0 {
            target, isImported := formatSysImport(ref)
            if isImported {
                g.addEdgeIfNotExist(Edge{
                    sourceId: id,
                    targetId: target,
                })
            }
            continue
        }
        for _, file := range files {
            fileId := path.Join(file.Dir(), file.Name())
            g.addNodeIfNotExist(Node{
                id: fileId,
                label: fileId,
                parentId: file.Dir(),
            })
            g.addEdgeIfNotExist(Edge{
                sourceId: id,
                targetId: fileId,
            })
        }
    }
    return g, nil
}

func formatSysImport(ref db.ImportRef) (string, bool){
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

