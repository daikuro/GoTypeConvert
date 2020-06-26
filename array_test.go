package typeconv

import "fmt"

func ExampleToArray() {
	v := []interface{}{"aaa"}
	a := ToArray(v)
	fmt.Println(a.A)
	// Output:
	// [aaa]
}

func ExampleToArrayNil() {
	a := ToArray(nil)
	fmt.Println(a.A)
	// Output:
	// []
}
