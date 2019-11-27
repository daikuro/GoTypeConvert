package typeconv

import (
	"fmt"
)

func ExampleToFloat64() {
	f := ToFloat64(10)
	fmt.Println(f.A)
	// Output:
	// 10
}

func ExampleToFloat64Float64() {
	f := ToFloat64(10.0)
	fmt.Println(f.A)
	// Output:
	// 10
}

func ExampleToFloat64Float32() {
	f := ToFloat64(float32(10.0))
	fmt.Println(f.A)
	// Output:
	// 10
}

func ExampleToFloat64Default() {
	f := ToFloat64("", 0)
	fmt.Println(f.A)
	// Output:
	// 0
}
