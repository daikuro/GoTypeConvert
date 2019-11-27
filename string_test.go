package typeconv

import (
	"fmt"
)

func ExampleString() {
	fmt.Println(ToString("1").A)
	// Output:
	// 1
}

func ExampleInt() {
	fmt.Println(ToString(1).A)
	fmt.Println(ToString(1).IsNil)
	// Output:
	// 1
	// false
}

func ExampleNil() {
	fmt.Println(ToString(nil, "A").A)
	fmt.Println(ToString(nil, "A").IsNil)
	// Output:
	// A
	// true
}

func ExampleEmpty() {
	fmt.Println(ToString("", "A").A)
	fmt.Println(ToString("", "A").IsNil)
	// Output:
	//
	// false
}
