package Chart

const MaxSize = 7

// 邻接矩阵存存储结构
// 顶点类型
type VertexType struct {
	No   int    //记录顶点的编号
	Info string //记录顶点的信息
}

// 图的邻接矩阵类型
type MGraph struct {
	Edges [MaxSize][MaxSize]int //邻接矩阵定义，如果是有权图，则将 int改为 float
	Node  int                   //顶点数
	Edge  int                   //边数
	Vex   [MaxSize]VertexType   //存放图的顶点信息
}

// 创建图
func CreateMgraph(m *MGraph) {

	// -1代表无路径可达，0代表自身到自身的距离
	arr := [MaxSize][MaxSize]int{
		{0, 4, 6, 6, -1, -1, -1},
		{-1, 0, 1, -1, 7, -1, -1},
		{-1, -1, 0, -1, 6, 4, -1},
		{-1, -1, 2, 0, -1, 5, -1},
		{-1, -1, -1, -1, 0, -1, 6},
		{-1, -1, -1, -1, 1, 0, 8},
		{-1, -1, -1, -1, -1, -1, 0},
	}

	m.Node = MaxSize
	for i := 0; i < MaxSize; i++ {
		m.Vex[i].No = i
		m.Vex[i].Info = "node info"
		for j := 0; j < MaxSize; j++ {
			m.Edges[i][j] = arr[i][i]
			if arr[i][j] == 0 || arr[i][j] == -1 {
				continue
			}
			m.Edge++
		}
	}
}
