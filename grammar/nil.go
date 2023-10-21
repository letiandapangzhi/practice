package grammar

import "fmt"

func nilTest() {
	a := new([]string)
	b := new(map[string]string)
	c := new([]string)
	fmt.Println(a, b, c, *a, *b, *c, a == nil, b == nil, a == c)

	var d []string
	var e map[string]string
	var f interface{}
	f = d
	fmt.Println(d, e, d == nil, e == nil, f == nil)

	var i *int
	var j interface{}
	fmt.Println(i, j, i == nil, j == nil, i == j)

}

func nil_test() {
	nilTest()
}
