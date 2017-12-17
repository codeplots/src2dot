package graph

type Node struct {
    id          string
    label       string
    parentId    string
}

type Edge struct {
    sourceId    string
    targetId    string
}

type Graph struct {
    nodes       []Node
    edges       []Edge
}

func (g *Graph) GetNode(id string) (Node, bool) {
    for _, n := range (*g).nodes {
        if id == n.id {
            return n, true
        }
    }
    return Node{}, false
}

func (g *Graph) GetEdge(sourceId string, targetId string) (Edge, bool) {
    for _, e := range (*g).edges {
        if e.sourceId == sourceId && e.targetId == targetId {
            return e, true
        }
    }
    return Edge{}, false
}


func (g *Graph) addNodeIfNotExist(n Node) bool {
    _, found := (*g).GetNode(n.id)
    if !found {
        (*g).nodes = append((*g).nodes, n)
    }
    return !found
}

func (g *Graph) addEdgeIfNotExist(e Edge) bool {
    _, found := (*g).GetEdge(e.sourceId, e.targetId)
    if !found {
        (*g).edges = append((*g).edges, e)
    }
    return !found
}
