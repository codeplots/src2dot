package graph

import (
    "path"
    db "github.com/codeplots/src2dot/database"
)

func DependencyGraph(store db.Database) (Graph, error) {
    g := Graph {
        nodes: []Node{},
        edges: []Edge{},
        isDirected: true,
    }

    for _, ref := range store.GetImports() {
        id := path.Join(ref.Dir(), ref.Filename())
        g.addNodeIfNotExist(Node{
            id: id,
            label: ref.Filename(),
            parentId: ref.Dir(),
        })

        files, _ := store.GetImportedFiles(ref)
        if len(files) == 0 {
            g.edges = append(g.edges, Edge{
                sourceId: id,
                targetId: ref.Symbol(),
            })
            continue
        }
        for _, file := range files {
            fileId := path.Join(file.Dir(), file.Name())
            g.addNodeIfNotExist(Node{
                id: fileId,
                label: file.Name(),
                parentId: file.Dir(),
            })
            g.edges = append(g.edges, Edge{
                sourceId: id,
                targetId: fileId,
            })
        }
    }
    return g, nil
}
