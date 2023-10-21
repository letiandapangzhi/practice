package grammar

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student // nil
	return stu
	// People = stu
}

func interface_test() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		// People为iface，只是data为nil，整体不等于nil
		fmt.Println("BBBBBBB")
	}
	// BBBBBBB

	var iTest People
	if iTest == nil {
		fmt.Println("CCCCCC")
	}
}
