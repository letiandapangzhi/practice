package grammar

import "fmt"

func testMap(data map[string]string) {
	data["test"] = "test"
	fmt.Printf("%p,%v\n", &data, data)
}

func map_test() {
	a := make(map[string]string)
	testMap(a)
	fmt.Printf("main: %p,%v\n", &a, a)
}
