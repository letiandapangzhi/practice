package helloAlgorithm

import (
	"fmt"
	"log"
)

type graphAdjMat struct {
	vertices []int   // 顶点列表 元素为顶点值 索引为顶点索引
	adjMat   [][]int // 邻接矩阵 行列索引为顶点索引
}

// 获取顶点数量
func (g *graphAdjMat) size() int {
	return len(g.vertices)
}

// 添加顶点
func (g *graphAdjMat) addVertex(val int) {
	// 邻接矩阵新增一行 先获取顶点长度 初始化一行
	n := g.size()
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)
	// 遍历每一行新增一列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
	// 顶点列表新增顶点
	g.vertices = append(g.vertices, val)
}

// 删除顶点
func (g *graphAdjMat) removeVertex(index int) {
	if index > g.size() {
		return
	}
	// 删除顶点列表指定索引
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	// 删除邻接矩阵指定行索引
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	// 删除邻接矩阵指定列索引
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

// 添加边
func (g *graphAdjMat) addEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		log.Println("存在越界")
		return
	}
	// 无向图中，邻接矩阵关于主对角线对称
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

// 删除边
func (g *graphAdjMat) removeEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		log.Println("存在越界")
		return
	}
	// 无向图中，邻接矩阵关于主对角线对称
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

// 输出图
func (g *graphAdjMat) print() {
	fmt.Printf("顶点列表:%v \n", g.vertices)
	fmt.Printf("邻接矩阵:\n")
	for i := range g.adjMat {
		fmt.Printf("%v\n", g.adjMat[i])
	}
}

func NewGraphAdjMat(vertices []int, edges [][]int) *graphAdjMat {
	// 1,2,3->索引0,1,2 假定1与2，2与3连接 edges[[0,1],[1,2]]
	// 初始化邻接矩阵
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	// 初始化图
	g := &graphAdjMat{
		vertices: vertices,
		adjMat:   adjMat,
	}
	// 添加边
	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])
	}
	return g
}
