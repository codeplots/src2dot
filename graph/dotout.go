package graph

func (g *Graph) ToDot() string {
	dot := ""
	dot += "digraph {\n"
	if (*g).kind == CLASS_DIAGRAM {
		dot += "\trankdir=\"BT\"\n"
		dot += "\tnode [shape=\"record\"]\n"
		dot += "\n"
	}
	for _, n := range (*g).nodes {
		dot += "\t\"" + n.id + "\" [label=\"" + n.label + "\"]\n"
	}
	dot += "\n"
	for _, e := range (*g).edges {
		switch e.style {
		case LINE:
			dot += "\t\"" + e.sourceId + "\" -> \"" + e.targetId + "\" [dir=none]\n"
		case HOLLOW_DIAMOND:
			dot += "\t\"" + e.sourceId + "\" -> \"" + e.targetId + "\" [arrowhead=\"ediamond\"]\n"
		case FILLED_DIAMOND:
			dot += "\t\"" + e.sourceId + "\" -> \"" + e.targetId + "\" [arrowhead=\"diamond\"]\n"
		default:
			dot += "\t\"" + e.sourceId + "\" -> \"" + e.targetId + "\"\n"
		}
	}
	dot += "}\n"
	return dot
}
