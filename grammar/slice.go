package grammar

import "fmt"

func slice_tests() {
	a := []int{1, 2, 3, 4}
	b := a[:1] // [0:1]
	c := a[1:] // [1:4]
	fmt.Printf("a len:%d cap:%d\n", len(a), cap(a))
	fmt.Printf("b len:%d cap:%d\n", len(b), cap(b)) // 0-1=1 4-0=4
	fmt.Printf("c len:%d cap:%d\n", len(c), cap(c)) // 4-1=1 4-0=4
}
