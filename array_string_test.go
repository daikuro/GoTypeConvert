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
