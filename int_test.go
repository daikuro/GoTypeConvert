package typeconv

import (
	"fmt"
)

func ExampleInt0() {
	fmt.Println(ToInt(1).A)
	// Output:
	// 1
}

func ExampleInt8() {
	fmt.Println(ToInt(int8((1))).A)
	// Output:
	// 1
}

func ExampleInt16() {
	fmt.Println(ToInt(int16((1))).A)
	// Output:
	// 1
}

func ExampleInt32() {
	fmt.Println(ToInt(int32((1))).A)
	// Output:
	// 1
}

func ExampleInt64() {
	fmt.Println(ToInt(int64((1))).A)
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

func ExampleIntFloat64() {
	fmt.Println(ToInt(1.1, 0).A)
	// Output:
	// 1
}

func ExampleIntFloat32() {
	fmt.Println(ToInt(float32(1.1), 0).A)
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
