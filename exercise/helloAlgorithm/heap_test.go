package helloAlgorithm

import (
	"container/heap"
	"fmt"
	"testing"
)

// 通过实现heap接口实现大顶堆
// 实现heap需要同时实现sort接口
type IntHead []int

// heap接口
func (t *IntHead) Push(val any) {
	*t = append(*t, val.(int))
}

func (t *IntHead) Pop() any {
	last := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return last
}

// sort接口
func (t *IntHead) Len() int {
	return len(*t)
}

func (t *IntHead) Less(i, j int) bool {
	return (*t)[i] > (*t)[j]
}

func (t *IntHead) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

// 获取堆顶
func (t *IntHead) Top() int {
	return (*t)[0]
}

type MyHeap struct {
	Data []int
}

func (t *MyHeap) left(i int) int {
	return 2*i + 1
}

func (t *MyHeap) right(i int) int {
	return 2*i + 2
}

func (t *MyHeap) parent(i int) int {
	return (i - 1) / 2
}

func (t *MyHeap) siftUp(i int) {
	for {
		parent := t.parent(i)
		// 越过根节点或子节点不大于父节点，结束堆化
		if parent < 0 || t.Data[i] <= t.Data[parent] {
			break
		}
		// 交换
		t.Data[parent], t.Data[i] = t.Data[i], t.Data[parent]
		// 继续向上堆化
		i = parent
	}
}

func (t *MyHeap) Push(val int) {
	t.Data = append(t.Data, val)
	t.siftUp(len(t.Data) - 1)
}

func (t *MyHeap) siftDown(i int) {
	for {
		// 获取父左右节点最大
		left := t.left(i)
		right := t.right(i)
		max := i
		if left < len(t.Data) && t.Data[left] > t.Data[max] {
			max = left
		}
		if right < len(t.Data) && t.Data[right] > t.Data[max] {
			max = right
		}
		// 最大还是父节点或者左右节点越界，不需堆化
		if max == i {
			break
		}
		// 交换
		t.Data[i], t.Data[max] = t.Data[max], t.Data[i]
		// 继续向下堆化
		i = max
	}
}

func (t *MyHeap) Pop() int {
	if len(t.Data) == 0 {
		fmt.Println("堆不存在数据")
		return -1
	}
	// 交换堆顶与堆底
	t.Data[0], t.Data[len(t.Data)-1] = t.Data[len(t.Data)-1], t.Data[0]
	// 删除现堆底(原堆顶)
	top := t.Data[len(t.Data)-1]
	t.Data = t.Data[:len(t.Data)-1]
	// 从顶部堆化
	t.siftDown(0)

	return top
}

func NewMyHeap(data []int) *MyHeap {
	myHeap := MyHeap{data}
	// 从最后一个父节点开始倒序的向下堆化->以该节点为根节点的子树都是都是合法的子堆
	for i := myHeap.parent(len(myHeap.Data) - 1); i >= 0; i-- {
		myHeap.siftDown(i)
	}

	return &myHeap
}

func TestHeap(t *testing.T) {
	// 初始化大顶堆
	maxHead := new(IntHead)
	heap.Init(maxHead)
	// 入堆
	heap.Push(maxHead, 1)
	heap.Push(maxHead, 2)
	heap.Push(maxHead, 3)
	heap.Push(maxHead, 4)
	// 堆顶
	fmt.Println("堆顶:", maxHead.Top())
	// 出堆
	fmt.Println("出堆:", heap.Pop(maxHead))
	fmt.Println("出堆:", heap.Pop(maxHead))
	fmt.Println("出堆:", heap.Pop(maxHead))
	fmt.Println("出堆:", heap.Pop(maxHead))

	data := []int{0, 1, 2, 3, 4, 5}
	myHeap := NewMyHeap(data)
	myHeap.Push(6)
	fmt.Println(myHeap.Data)
	fmt.Println("出堆:", myHeap.Pop())
	fmt.Println("出堆:", myHeap.Pop())
	fmt.Println("出堆:", myHeap.Pop())
	fmt.Println("出堆:", myHeap.Pop())
}
