package graph

type Node struct {
    Id       string `json:"id"`
    Label    string `json:"label"`
    ParentId string `json:"parent"`
}

type EdgeStyle string

const (
	LINE           EdgeStyle = "line"
	ARROW          EdgeStyle = "arrow"
	HOLLOW_DIAMOND EdgeStyle = "hollow_diamond"
	FILLED_DIAMOND EdgeStyle = "filled_diamond"
)

type Edge struct {
    SourceId string `json:"source"`
    TargetId string `json:"target"`
    Style    EdgeStyle      `json:"style"`
}

type GraphType string

const (
	DEPENDENCY_GRAPH GraphType = "dependency_graph"
	CLASS_DIAGRAM    GraphType = "class_diagram"
	CALL_GRAPH       GraphType = "call_graph"
)

type Graph struct {
    Nodes []Node  `json:"nodes"`
    Edges []Edge `json:"edges"`
    Kind  GraphType `json:"kind"`
}

func (g *Graph) GetNode(id string) (Node, bool) {
	for _, n := range (*g).Nodes {
		if id == n.Id {
			return n, true
		}
	}
	return Node{}, false
}

func (g *Graph) GetEdge(sourceId string, targetId string) (Edge, bool) {
	for _, e := range (*g).Edges {
		if e.SourceId == sourceId && e.TargetId == targetId {
			return e, true
		}
	}
	return Edge{}, false
}

func (g *Graph) addNodeIfNotExist(n Node) bool {
	_, found := (*g).GetNode(n.Id)
	if !found {
		(*g).Nodes = append((*g).Nodes, n)
	}
	return !found
}

func (g *Graph) addEdgeIfNotExist(e Edge) bool {
	_, found := (*g).GetEdge(e.SourceId, e.TargetId)
	if !found {
		(*g).Edges = append((*g).Edges, e)
	}
	return !found
}
