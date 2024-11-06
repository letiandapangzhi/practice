package helloAlgorithm

import (
	"container/list"
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	// 初始化
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)
	n6 := NewTreeNode(6)
	n7 := NewTreeNode(7)

	n1.Left = n2
	n1.Right = n3

	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	// 遍历输出
	n1.Print()

	// 二叉搜索树
	search := NewSearchTree([]int{1, 2, 3})
	// 遍历输出
	search.Root.Print()
}

// 二叉树节点结构体
type TreeNode struct {
	Val   int       // 节点值
	Left  *TreeNode // 左子节点
	Right *TreeNode // 右子节点
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (t *TreeNode) Print() {
	// 层序遍历
	t.LevelOrder()
	// 前序遍历
	fmt.Println("前序遍历")
	t.PrevOrder()
	fmt.Println()
	// 中序遍历
	fmt.Println("中序遍历")
	t.InOrder()
	fmt.Println()
	// 后序遍历
	fmt.Println("后序遍历")
	t.NextOrder()
	fmt.Println()
}

// 层序遍历
func (t *TreeNode) LevelOrder() {
	fmt.Println("层序遍历")
	queue := list.New()
	// 节点入队列
	queue.PushBack(t)

	// 遍历队列
	for queue.Len() > 0 {
		// 节点出队列
		node := queue.Remove(queue.Front()).(*TreeNode)
		fmt.Printf("%d", node.Val)

		if node.Left != nil {
			// 左子节点入队
			queue.PushBack(node.Left)
			fmt.Printf("->%d", node.Left.Val)
		}
		if node.Right != nil {
			// 右子节点入队
			queue.PushBack(node.Right)
			fmt.Printf("->%d", node.Right.Val)
		}
		fmt.Println()
	}
}

// 前序遍历 递归
func (t *TreeNode) PrevOrder() {
	if t == nil {
		return
	}
	fmt.Printf("%d ", t.Val)
	t.Left.PrevOrder()
	t.Right.PrevOrder()
}

// 中序遍历 递归
func (t *TreeNode) InOrder() {
	if t == nil {
		return
	}
	t.Left.InOrder()
	fmt.Printf("%d ", t.Val)
	t.Right.InOrder()
}

// 后序遍历 递归
func (t *TreeNode) NextOrder() {
	if t == nil {
		return
	}
	t.Left.NextOrder()
	t.Right.NextOrder()
	fmt.Printf("%d ", t.Val)
}

// 二叉搜索树
type SearchTree struct {
	Root *TreeNode // 根节点指针
}

func NewSearchTree(data []int) *SearchTree {
	fmt.Println("构建二叉搜索树", data)
	node := new(SearchTree)
	for _, val := range data {
		node.Insert(val)
	}
	return node
}

// 插入
func (s *SearchTree) Insert(val int) {
	// 树根节点
	node := s.Root
	// 初始化树
	if node == nil {
		s.Root = NewTreeNode(val)
		return
	}

	// 遍历直到越过叶子节点
	prev := new(TreeNode)
	for node != nil {
		// 去重
		if node.Val == val {
			return
		}
		// 记录越过叶子节点时的节点
		prev = node
		// 对比确认分支，节点挪移
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	// 从记录的节点判断分支插入新节点
	newNode := NewTreeNode(val)
	if val < prev.Val {
		prev.Left = newNode
	} else {
		prev.Right = newNode
	}

}
