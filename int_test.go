package typeconv

import (
	"fmt"
)

func ExampleInt0() {
	fmt.Println(ToInt(1).A)
	// Output:
	// 1
}

func ExampleInt1() {
	fmt.Println(ToInt(1, 0).A)
	// Output:
	// 1
}

func ExampleInt2() {
	fmt.Println(ToInt("2", 0).A)
	// Output:
	// 2
}

func ExampleInt3() {
	fmt.Println(ToInt(-1, 0).A)
	// Output:
	// -1
}

func ExampleInt4() {
	fmt.Println(ToInt(1.1, 0).A)
	// Output:
	// 1
}

func ExampleInt5() {
	fmt.Println(ToInt("e", 0).A)
	fmt.Println(ToInt("e", 0).Error)
	// Output:
	// 0
	// strconv.Atoi: parsing "e": invalid syntax
}
func ExampleBool() {
	fmt.Println(ToInt(true, 100).A)
	fmt.Println(ToInt(false, 100).A)
	// Output:
	// 1
	// 0
}
