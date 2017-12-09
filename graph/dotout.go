package graph

func (g *Graph) ToDot() string {
    dot := ""
    switch (*g).isDirected {
    case true:
        dot += "digraph G {\n"
    case false:
        dot += "graph G {\n"
    }
    for _, e := range (*g).edges {
        dot += "\t\"" + e.sourceId + "\" -> \"" + e.targetId + "\"\n"
    }
    dot += "}\n"
    return dot
}
