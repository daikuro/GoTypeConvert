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

func ExampleToFloat64Int8() {
	f := ToFloat64(int8(10))
	fmt.Println(f.A)
	// Output:
	// 10
}

func ExampleToFloat64Int16() {
	f := ToFloat64(int16(10))
	fmt.Println(f.A)
	// Output:
	// 10
}

func ExampleToFloat64Int32() {
	f := ToFloat64(int32(10))
	fmt.Println(f.A)
	// Output:
	// 10
}

func ExampleToFloat64Int64() {
	f := ToFloat64(int32(10))
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
