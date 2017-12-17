package graph

func (g *Graph) ToDot() string {
    dot := ""
    dot += "digraph {\n"
    }
    for _, n := range (*g).nodes {
        dot += "\t\"" + n.id + "\" [label=\"" + n.label + "\"]\n"
    }
    dot += "\n"
    for _, e := range (*g).edges {
        dot += "\t\"" + e.sourceId + "\" -> \"" + e.targetId + "\"\n"
    }
    dot += "}\n"
    return dot
}
