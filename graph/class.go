package graph

import (
    db "github.com/codeplots/src2dot/database"
)

func ClassDiagram(store db.Database) (Graph, error) {
    g := Graph{
        kind: CLASS_DIAGRAM,
    }
    classes := store.GetClasses()
    for _, c := range classes {
        label := "{"
        label += c.Symbol()
        label += "|"

        for _, m := range store.GetMembers(c) {
            label += m.Symbol() 
            if m.Typeref() != "" {
                label += " : " + m.Typeref()
            }
            label += "\\l"
        }
        label += "|"
        for _, m := range store.GetMethods(c) {
            label += m.Symbol() + m.Signature()
            if m.Typeref() != "" {
                label += " : " + m.Typeref()
            }
            label += "\\l"
        }
        label += "}"

        g.addNodeIfNotExist(Node{
            id: c.Symbol(),
            label: label,
        })

        for _, p := range store.GetParents(c) {
            g.addEdgeIfNotExist(Edge{
                sourceId: c.Symbol(),
                targetId: p.Symbol(),
                style: ARROW,
            })
        }
        for _, a := range store.GetAssociates(c) {
            g.addEdgeIfNotExist(Edge{
                sourceId: c.Symbol(),
                targetId: a.Symbol(),
                style: LINE,
            })
        }
        for _, a := range store.GetAggregations(c) {
            g.addEdgeIfNotExist(Edge{
                sourceId: a.Symbol(),
                targetId: c.Symbol(),
                style: HOLLOW_DIAMOND,
            })
        }
    }
    return g, nil
}
