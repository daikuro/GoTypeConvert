package typeconv

import "fmt"

func ExampleBoolString1() {
	fmt.Println(ToBool("true").A)
	// Output:
	// true
}
func ExampleBoolString2() {
	fmt.Println(ToBool("false").A)
	// Output:
	// false
}
func ExampleBoolBool1() {
	fmt.Println(ToBool(true).A)
	// Output:
	// true
}
func ExampleBoolBool2() {
	fmt.Println(ToBool(false).A)
	// Output:
	// false
}
func ExampleBoolInt1() {
	fmt.Println(ToBool(1).A)
	// Output:
	// true
}
func ExampleBoolInt0() {
	fmt.Println(ToBool(0).A)
	// Output:
	// false
}
func ExampleBoolFloat64_1() {
	fmt.Println(ToBool(1.0).A)
	// Output:
	// true
}
func ExampleBoolFloat64_0() {
	fmt.Println(ToBool(0.0).A)
	// Output:
	// false
}
func ExampleBoolDefaultValueTrue() {
	fmt.Println(ToBool(nil, true).A)
	// Output:
	// true
}
func ExampleBoolDefaultValuefalse() {
	fmt.Println(ToBool(nil, false).A)
	// Output:
	// false
}
