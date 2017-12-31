package graph

func (g *Graph) ToDot() string {
	dot := ""
	dot += "digraph {\n"
	if (*g).Kind == CLASS_DIAGRAM {
		dot += "\trankdir=\"BT\"\n"
		dot += "\tnode [shape=\"record\"]\n"
		dot += "\n"
	}
	for _, n := range (*g).Nodes {
		dot += "\t\"" + n.Id + "\" [label=\"" + n.Label + "\"]\n"
	}
	dot += "\n"
	for _, e := range (*g).Edges {
		switch e.Style {
		case LINE:
			dot += "\t\"" + e.SourceId + "\" -> \"" + e.TargetId + "\" [dir=none]\n"
		case HOLLOW_DIAMOND:
			dot += "\t\"" + e.SourceId + "\" -> \"" + e.TargetId + "\" [arrowhead=\"ediamond\"]\n"
		case FILLED_DIAMOND:
			dot += "\t\"" + e.SourceId + "\" -> \"" + e.TargetId + "\" [arrowhead=\"diamond\"]\n"
		default:
			dot += "\t\"" + e.SourceId + "\" -> \"" + e.TargetId + "\"\n"
		}
	}
	dot += "}\n"
	return dot
}
