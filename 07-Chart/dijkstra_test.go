package Chart

import "testing"

func TestDijkstra(T *testing.T) {

	dist := make([]int, MaxSize)
	path := make([]int, MaxSize)

	g := &MGraph{}
	CreateMgraph(g)

	Dijkstra(g, 3, dist, path)
}
