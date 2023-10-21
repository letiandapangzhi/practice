package grammar

type T string

func (t *T) echo() {
	println("hello")
}

func funcTest() {
	var t1 T = "ABC"
	t1.echo()
	// 编译不通过
	// map对象中的元素无法寻址
	//a := map[string]T{
	//	"test": "ABC",
	//}
	//a["test"].echo()
	// 常量无法寻址
	//const b T = "ABC"
	//b.echo()
}

func func_test() {
	funcTest()
}
