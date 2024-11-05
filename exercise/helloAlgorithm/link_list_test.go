package main

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	// 初始化链表节点
	n0 := NewListNode(0)
	n1 := NewListNode(1)
	n2 := NewListNode(2)
	n3 := NewListNode(3)

	// 链接
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3

	// 输出
	n0.Print()

	// 在n1后插入节点4
	n4 := NewListNode(4)
	n1.Insert(n4)

	// 输出
	n0.Print()

	// 删除节点n1后的节点
	n1.Delete(n2)
	n1.Delete(n4)

	// 输出
	n0.Print()

	// 查找第5个节点
	n0.Index(4)
	// 查找第4个节点
	n0.Index(3)

	// 查找值10
	n0.Find(10)
	// 查找值2
	n0.Find(2)

}

// 链表节点结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 下一个节点地址
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 链表输出
func (l *ListNode) Print() {
	fmt.Printf("%d", l.Val)
	node := l.Next
	num := 0
	for node != nil {
		fmt.Printf(" -> %d", node.Val)
		node = node.Next
		num++
		if num > 100 {
			break
		}
	}
	fmt.Println()
}

// 在链表节点之后插入节点
func (l *ListNode) Insert(node *ListNode) {
	node.Next = l.Next
	l.Next = node
}

// 删除节点之后的节点
func (l *ListNode) Delete(node *ListNode) {
	if l.Next != node {
		fmt.Printf("被删除节点%d不是节点%d的下一个节点\n", node.Val, l.Val)
		return
	}
	l.Next = node.Next
	node.Next = nil
}

// 检索链表的第几个索引
func (l *ListNode) Index(index int) {
	node := l
	for i := 0; i < index; i++ {
		if node == nil || node.Next == nil {
			fmt.Printf("索引%d不存在\n", index)
			return
		}
		node = node.Next
	}
	fmt.Printf("索引%d对应节点%d\n", index, node.Val)
}

// 查找链表是否存在某个值
func (l *ListNode) Find(val int) {
	node := l
	index := 0
	for node.Next != nil {
		if node.Val == val {
			fmt.Printf("值%d首次出现在链表的索引%d位置", val, index)
			return
		}
		index++
		node = node.Next
	}
	fmt.Printf("值%d不存在链表里\n", val)
}
