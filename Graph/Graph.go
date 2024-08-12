package main

import (
	"fmt"
)

type Graph struct {
	vertices []int
	edges    [][]int
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(v int) {
	g.vertices = append(g.vertices, v)
	g.edges = append(g.edges, []int{})
}

func (g *Graph) AddEdge(src, dst int) {
	g.edges[src] = append(g.edges[src], dst)
}

func (g *Graph) Print() {
	fmt.Println("Vertices: ", g.vertices)
	fmt.Println("Edges: ", g.edges)
}

func main() {
	g := NewGraph()

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)

	g.AddEdge(0, 1)
	g.AddEdge(1, 0)
	g.AddEdge(1, 2)

	g.Print()
}
