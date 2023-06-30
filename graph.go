package main

import (
	"fmt"
)

type Vertex struct {
	id       int
	Vertices []*Vertex
}

func NewVertex(id int) *Vertex {
	return &Vertex{id: id}
}

type Graph struct {
	Vertices []*Vertex
}

// get vertex

func (g *Graph) GetVertex(vertex_id int) *Vertex {
	for _, v := range g.Vertices {
		if v.id == vertex_id {
			return v
		}
	}
	return nil
}

// add vertex
func (g *Graph) AddVertex(id int) {
	g.Vertices = append(g.Vertices, NewVertex(id))
}

// add edge
func (g *Graph) AddEdge(a_id int, b_id int) {
	A := g.GetVertex(a_id)
	B := g.GetVertex(b_id)

	if A == nil || B == nil {
		return
	}

	A.Vertices = append(A.Vertices, B)
	B.Vertices = append(B.Vertices, A)
}

// Print Graph

func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("Vertex %d : ", v.id)
		for _, adj := range v.Vertices {
			fmt.Printf("%d ", adj.id)
		}
		fmt.Println("")
	}
}

func (g *Graph) BFS(vertex_id int) {
	// Start := g.GetVertex(vertex_id)

}

func (g *Graph) DFS() {

}

func (g *Graph) Dijkstra() {

}
