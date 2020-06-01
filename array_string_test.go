package typeconv

import "fmt"

func ExampleToArrayString() {
	a := ToArrayString([]string{"aaa"})
	fmt.Println(a.A)
	fmt.Println(a.T)
	// Output:
	// [aaa]
	// []string
}

func ExampleToArrayString2() {
	a := ToArrayString([]interface{}{"aaa", "2"})
	fmt.Println(a.A)
	fmt.Println(a.T)
	// Output:
	// [aaa 2]
	// []interface {}
}

func ExampleToArrayStringDefaultValue() {
	a := ToArrayString(nil, []string{"A"})
	fmt.Println(a.A)
	// Output:
	// [A]
}

func ExampleToArrayStringInt() {
	a := ToArrayString([]int{1})
	fmt.Println(a.A)
	// Output:
	// [1]
}

func ExampleToArrayStringInt8() {
	a := ToArrayString([]int8{1})
	fmt.Println(a.A)
	// Output:
	// [1]
}

func ExampleToArrayStringInt32() {
	a := ToArrayString([]int32{1})
	fmt.Println(a.A)
	// Output:
	// [1]
}

func ExampleToArrayStringInt64() {
	a := ToArrayString([]int64{1})
	fmt.Println(a.A)
	// Output:
	// [1]
}

func ExampleToArrayStringFloat32() {
	a := ToArrayString([]float32{1})
	fmt.Println(a.A)
	// Output:
	// [1]
}

func ExampleToArrayStringFloat64() {
	a := ToArrayString([]float64{1})
	fmt.Println(a.A)
	// Output:
	// [1]
}

func ExampleToArrayStringBool() {
	a := ToArrayString([]bool{true, false})
	fmt.Println(a.A)
	// Output:
	// [true false]
}

func ExampleArrayStringContains() {
	a := []string{"aaa", "bbb", "ccc"}
    r := ArrayStringContains(a, "bbb")
    fmt.Println(r)
    // Output:
    // true
}

func ExampleArrayStringContainsFalse() {
	a := []string{"aaa", "bbb", "ccc"}
    r := ArrayStringContains(a, "ddd")
    fmt.Println(r)
    // Output:
    // false
}