package exercise

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	var add int
	if l1.Val+l2.Val > 9 {
		l3.Val = l1.Val + l2.Val - 10
		add = 1
	} else {
		l3.Val = l1.Val + l2.Val
	}
	// 游标
	ran1 := l1.Next
	ran2 := l2.Next
	// 临时保留上一个节点
	prev := new(ListNode)

	first := true
	for ran1 != nil || add != 0 {
		if ran1 == nil {
			ran1 = new(ListNode)
		}
		if ran2 == nil {
			ran2 = new(ListNode)
		}

		node := new(ListNode)
		if ran1.Val+ran2.Val+add > 9 {
			node.Val = ran1.Val + ran2.Val + add - 10
			add = 1
		} else {
			node.Val = ran1.Val + ran2.Val + add
			add = 0
		}
		if first {
			l3.Next = node
			first = false
		}
		ran1 = ran1.Next
		ran2 = ran2.Next

		prev.Next = node
		prev = node

	}
	return l3
}

func addTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	// 尾部，指向下一个节点，起游标作用
	var tail *ListNode
	// 进值
	var carry int

	for l1 != nil || l2 != nil {
		// 数值
		var n1, n2 int

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}
