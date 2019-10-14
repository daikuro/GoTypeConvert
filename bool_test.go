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
