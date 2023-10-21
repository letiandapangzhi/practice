package grammar

import "fmt"

func makeNewTest() {
	a := make([]int, 0, 5)
	b := new([]int)
	fmt.Println(a, *b, a == nil, *b == nil)

	c := make(map[string]string)
	d := new(map[string]string)
	fmt.Println(c, *d, c == nil, *d == nil)
}

func make_new_test() {
	makeNewTest()
}
