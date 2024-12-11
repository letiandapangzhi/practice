package helloAlgorithm

import (
	"container/list"
	"fmt"
	"testing"
)

func TestBalanceTree(t *testing.T) {
	bTree := NewBalanceTree()
	bTree.Insert(1)
	bTree.Insert(2)
	bTree.Insert(3)
	bTree.Insert(4)
	bTree.Root.Print()
	bTree.Remove(3)
	bTree.Root.Print()
}

// 平衡二叉搜索树节点结构体
type BalanceTreeNode struct {
	Val    int              // 节点值
	Left   *BalanceTreeNode // 左子节点
	Right  *BalanceTreeNode // 右子节点
	height int              // 节点高度
}

// 构建树节点
func NewBalanceTreeNode(val int) *BalanceTreeNode {
	return &BalanceTreeNode{
		Val:    val,
		Left:   nil,
		Right:  nil,
		height: 0,
	}
}

// 打印树
func (t *BalanceTreeNode) Print() {
	// 层序遍历
	t.LevelOrder()
}

// 层序遍历
func (t *BalanceTreeNode) LevelOrder() {
	fmt.Println("层序遍历")
	queue := list.New()
	// 节点入队列
	queue.PushBack(t)

	// 遍历队列
	for queue.Len() > 0 {
		// 节点出队列
		node := queue.Remove(queue.Front()).(*BalanceTreeNode)
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

// 获取节点高度
func (b *BalanceTreeNode) GetHeight() int {
	if b != nil {
		return b.height
	}
	return -1
}

// 更新节点高度
func (b *BalanceTreeNode) UpdateHeight() {
	// 获取左右子节点高度
	lh := b.Left.GetHeight()
	rh := b.Right.GetHeight()
	// 节点高度=最高子树高度 + 1
	if lh > rh {
		b.height = lh + 1
	} else {
		b.height = rh + 1
	}
}

// 获取平衡因子 左子树与右子树高度差
func (b *BalanceTreeNode) GetBalanceFactor() int {
	if b == nil {
		return 0
	}
	return b.Left.GetHeight() - b.Right.GetHeight()
}

// 节点右旋 以失衡节点的左子节点为原点将失衡节点旋转到左子节点的右子节点 父->子的子->孙 🤣
// 左链表 3 2 1，3失衡
func (b *BalanceTreeNode) RightRotate() *BalanceTreeNode {
	// 失衡节点[b]的左子节点
	lChild := b.Left
	// 左子节点的右子节点 失衡节点的孙子 可能不存在
	rGrandChild := lChild.Right
	// 失衡节点右旋成了左子节点的右子节点，替换掉了原来的右子节点
	lChild.Right = b
	// 原来的右子节点要变成失衡节点的左子节点
	b.Left = rGrandChild
	// 更新层级变动节点的高度 左子节点->父 失衡节点->子
	// 变成子节点的失衡节点需要先更新 父节点有依赖 原右子节点的左右子树不变，高度不变
	b.UpdateHeight()
	lChild.UpdateHeight()
	// 返回新的子树根节点
	return lChild
}

// 节点左旋 以失衡节点的右子节点为原点将失衡节点旋转到右子节点的左子节点
// 右链表 1 2 3，1失衡
func (b *BalanceTreeNode) LeftRotate() *BalanceTreeNode {
	// 失衡节点[b]的右子节点
	rChild := b.Right
	// 右子节点的左子节点
	lGrandChild := rChild.Left
	// 失衡节点左旋到右子节点的左子节点
	rChild.Left = b
	// 原左子节点变成失衡节点的右子节点
	b.Right = lGrandChild
	// 更新高度
	b.UpdateHeight()
	rChild.UpdateHeight()
	// 返回左旋后子树的根节点
	return rChild
}

// 旋转 根据失衡节点的平衡因子与偏向侧子节点的平衡因子正负决定旋转
// 右旋=失衡节点右旋[3 2 1] 左旋=失衡节点左旋[1 2 3]
// 先左旋再右旋=子节点先左旋再失衡节点右旋[3 1 2]
// 先右旋再左旋=子节点先右旋再失衡节点左旋[1 3 2]
// 返回平衡后子树的根节点
func (b *BalanceTreeNode) Rotate() *BalanceTreeNode {
	// 获取节点平衡因子 范围[-2,-1,0,1,2] 绝对值大于1为失衡
	bf := b.GetBalanceFactor()
	if bf > 1 {
		// 左子树高，整体偏左，失衡节点需要右旋
		if b.Left.GetBalanceFactor() >= 0 {
			// 失衡节点的左子节点平衡因子正数->偏向左，无需左旋
			return b.RightRotate()
		} else {
			// 负数->偏向右，需要左旋让整体偏左配合父节点右旋
			b.Left = b.Left.LeftRotate()
			return b.RightRotate()
		}
	} else if bf < -1 {
		// 右子树高，整体偏右，失衡节点需要左旋
		if b.Right.GetBalanceFactor() >= 0 {
			// 失衡节点的右子节点平衡因子正数->偏向左，需要右旋让整体偏右配合父节点左旋
			b.Right = b.Right.RightRotate()
			return b.LeftRotate()
		} else {
			// 负数->偏向右与整体一致，父节点左旋
			return b.LeftRotate()
		}
	}
	// 平衡无需旋转
	return b
}

// 插入数据 返回子树根节点
func (b *BalanceTreeNode) Insert(val int) *BalanceTreeNode {
	// 空树初始化
	if b == nil {
		return NewBalanceTreeNode(val)
	}
	// 递归查找位置
	if val < b.Val {
		// 待插入值小于当前节点应该在节点的左子树 以左子树为根节点递归尝试插入 直到重复或者适合位置[空]创建新节点，返回节点指针
		b.Left = b.Left.Insert(val)
	} else if val > b.Val {
		// 待插入值大于当前节点应该在节点的右子树 以右子树为根节点递归尝试插入 直到重复或者适合位置[空]创建新节点，返回节点指针
		b.Right = b.Right.Insert(val)
	} else {
		// 重复无需插入
		return b
	}
	// 更新当前节点高度，当前节点的子节点插入新值 高度有变化
	b.UpdateHeight()
	// 再旋转，依赖高度
	return b.Rotate()
}

// 删除数据 返回子树根节点
func (b *BalanceTreeNode) Remove(val int) *BalanceTreeNode {
	if b == nil {
		return nil
	}
	if val < b.Val {
		b.Left = b.Left.Remove(val)
	} else if val > b.Val {
		b.Right = b.Right.Remove(val)
	} else {
		if b.Left == nil || b.Right == nil {
			child := b.Left
			if b.Right != nil {
				child = b.Right
			}
			if child == nil {
				// 子节点 = 0，删除节点直接返回
				return nil
			} else {
				// 子节点 = 1，子节点替换父节点
				b = child
			}
		} else {
			// 子节点 = 2，则将中序遍历的下个节点删除，并用该节点替换当前节点
			child := b.Right
			if child.Left != nil {
				child = child.Left
			}
			// 将中序遍历的下个节点删除
			b.Right = b.Right.Remove(child.Val)
			// 该节点替换当前节点
			b.Val = child.Val
		}
	}
	// 更新高度
	b.UpdateHeight()
	// 旋转平衡
	return b.Rotate()
}

// 二叉平衡搜索树
type BalanceTree struct {
	Root *BalanceTreeNode
}

func NewBalanceTree() *BalanceTree {
	return &BalanceTree{}
}

func (b *BalanceTree) Insert(val int) {
	b.Root = b.Root.Insert(val)
}

func (b *BalanceTree) Remove(val int) {
	b.Root = b.Root.Remove(val)
}
