package grammar

import (
	"fmt"
)

func slice_tests() {
	a := []int{1, 2, 3, 4}
	b := a[:1] // [0:1]
	c := a[1:] // [1:4]
	fmt.Printf("a len:%d cap:%d\n", len(a), cap(a))
	fmt.Printf("b len:%d cap:%d\n", len(b), cap(b)) // 0-1=1 4-0=4
	fmt.Printf("c len:%d cap:%d\n", len(c), cap(c)) // 4-1=1 4-0=4

	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice(s1)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d\n", s, len(s), cap(s))
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d\n", s1, len(s1), cap(s1))
	// s: [0 0 0 0 0 0 0 0 0 0], len of s: 10, cap of s: 12
	// s1: [0 0], len of s1: 2, cap of s1: 4

	a1 := make([]int, 0, 12)
	changeSlice(a1)
	fmt.Printf("a1: %v, len of a1: %d, cap of a1: %d\n", a1, len(a1), cap(a1))
	changeSlice1(&a1)
	fmt.Printf("a1: %v, len of a1: %d, cap of a1: %d\n", a1, len(a1), cap(a1))
}

func changeSlice(s1 []int) {
	s1 = append(s1, 10)
	fmt.Printf("func s1: %v, len of s1: %d, cap of s1: %d\n", s1, len(s1), cap(s1))
}

func changeSlice1(s1 *[]int) {
	*s1 = append(*s1, 10)
}
