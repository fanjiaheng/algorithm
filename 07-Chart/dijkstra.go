package Chart

// distp[vi]：表示当前已找到的从v0到每个终点vi的最短路径长度。
// 初态：若v0到vi有边，则distp[vi]为边上的权值，否则distp[vi]=-1

// path[vi]：表示保存v0到vi最短路径上vi的前一个顶点

// set[]：标记数组，set[vi]=0表示vi在T中，即没有被并入最短路径；
// set[vi]=1表示vi在S中，即已经被并入最短路径。初态：set[v0]=1,其余全为零。
func Dijkstra(graph *MGraph, v int, dist []int, path []int) {

	var set [MaxSize]int
	var u int

	// 初始化操作
	for i := 0; i < graph.Node; i++ {
		dist[i] = graph.Edges[v][i]
		set[i] = 0
		if graph.Edges[v][i] != -1 {
			path[i] = v
		} else {
			path[i] = -1
		}
	}
	set[v] = 1
	path[v] = -1

	for i := 0; i < graph.Node-1; i++ {
		min := 9999999

		for j := 0; j < graph.Node; j++ {
			if set[j] == 0 && dist[j] < min {
				u = j
				min = dist[j]
			}
		}

		set[u] = 1
		for j := 0; j < graph.Node; j++ {
			if set[j] == 0 && dist[u]+graph.Edges[u][j] < dist[j] {
				dist[j] = dist[u] + graph.Edges[u][j]
				path[j] = u
			}
		}
	}
}
