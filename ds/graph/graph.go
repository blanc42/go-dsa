package graph

import "fmt"

// edge struct (to -> from)
type toedge struct {
	to     int
	weight int
}

// Graph contains slice of all Vertices, count of those and Map of Edges
// unweighted graph
type Graph struct {
	Vertices []int
	Count    int
	Edges    map[int][]toedge
}

// returns a new graph struct
func NewGraph() *Graph {
	g := new(Graph)
	// You have to initialize the map using the make function (or a map literal) before you can add any elements:
	// https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/
	g.Edges = make(map[int][]toedge)
	return g
}

// vertex
func (g *Graph) AddVertex(v int) {
	g.Count++
	g.Vertices = append(g.Vertices, v)
}

func (g *Graph) AddEdge(f int, t int, w int) {
	e := toedge{t, w}
	g.Edges[f] = append(g.Edges[f], e)

}

func (g *Graph) GetVertex(int) {

}

func (g *Graph) Print() {
	fmt.Println("The graph has", g.Count, "vertex/vertices", g.Vertices, g.Edges)
	for _, v := range g.Vertices {
		fmt.Println(v, g.Edges[v])
	}
}
