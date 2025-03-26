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
	search := NewSearchTree([]int{2, 1, 4, 3, 5, 6})
	// 遍历输出
	search.Root.Print()
	// 搜索
	search.Find(4)
	search.Find(1)

	search.Delete(4)
	// 遍历输出
	search.Root.Print()
}

// 二叉树节点结构体
type TreeNode struct {
	Val   int       // 节点值
	Left  *TreeNode // 左子节点
	Right *TreeNode // 右子节点
}

// 构建树节点
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// 打印树
func (t *TreeNode) Print() {
	// 层序遍历
	t.LevelOrder()
	// // 前序遍历
	// fmt.Println("前序遍历")
	// t.PrevOrder()
	// fmt.Println()
	// // 中序遍历
	// fmt.Println("中序遍历")
	// t.InOrder()
	// fmt.Println()
	// // 后序遍历
	// fmt.Println("后序遍历")
	// t.NextOrder()
	// fmt.Println()
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

// 构建二叉搜索树
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

// 搜索
func (s *SearchTree) Find(val int) {
	node := s.Root
	if node == nil {
		fmt.Println("二叉搜索树不存在数据")
		return
	}

	for node != nil {
		if val == node.Val {
			fmt.Printf("%d 在二叉搜索树中存在\n", val)
			return
		} else if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	fmt.Printf("%d 在二叉搜索树中不存在\n", val)
}

// 删除
func (s *SearchTree) Delete(val int) {
	if s.Root == nil {
		fmt.Println("二叉搜索树不存在数据")
		return
	}

	// 记录删除节点的上一个节点
	prev := new(TreeNode)
	node := s.Root
	for node != nil {
		if node.Val == val {
			break
		}
		prev = node

		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	// 两种可能 找不到node = nil 找到 node 待删除节点
	if node == nil {
		fmt.Printf("%d 在二叉搜索树中不存在\n", val)
		return
	}
	// 待删除节点node
	if node.Left == nil || node.Right == nil {
		// node子节点数0或1
		// 获取删除节点子节点
		child := new(TreeNode)
		if node.Left != nil {
			child = node.Left
		} else {
			child = node.Right
		}

		if node == s.Root {
			// 删除节点是根节点
			s.Root = child
		} else {
			// 删除节点非根节点,判断删除节点是上一个节点的左还是右子节点
			if prev.Left == node {
				prev.Left = child
			} else {
				prev.Right = child
			}
		}
	} else {
		// node子节点数2
		// 从删除节点的右子树获取最小值，即右子树的中序遍历第一个节点
		tmp := node.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 递归删除第一个节点，避免它有右子树丢失
		s.Delete(tmp.Val)
		// 递归回来，替换删除节点值
		node.Val = tmp.Val
	}
}
