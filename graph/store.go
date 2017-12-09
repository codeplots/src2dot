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
    isDirected  bool
}

func (g *Graph) GetNode(id string) (Node, bool) {
    for _, n := range (*g).nodes {
        if id == n.id {
            return n, true
        }
    }
    return Node{}, false
}

func (g *Graph) addNodeIfNotExist(n Node) bool {
    _, found := (*g).GetNode(n.id)
    if !found {
        (*g).nodes = append((*g).nodes, n)
    }
    return !found
}
