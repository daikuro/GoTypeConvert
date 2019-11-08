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
