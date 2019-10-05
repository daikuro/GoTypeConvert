package typeconv

import (
	"fmt"
)

func ExampleInt1() {
	fmt.Println(ToIntd(1, 0).A)
	// Output:
	// 1
}

func ExampleInt2() {
	fmt.Println(ToIntd("2", 0).A)
	// Output:
	// 2
}

func ExampleInt3() {
	fmt.Println(ToIntd(-1, 0).A)
	// Output:
	// -1
}

func ExampleInt4() {
	fmt.Println(ToIntd(1.1, 0).A)
	// Output:
	// 1
}

func ExampleInt5() {
	fmt.Println(ToIntd("e", 0).A)
	fmt.Println(ToIntd("e", 0).Error)
	// Output:
	// 0
	// strconv.Atoi: parsing "e": invalid syntax
}
