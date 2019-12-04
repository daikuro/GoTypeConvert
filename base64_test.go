package typeconv

import "fmt"

func ExampleToBase64() {
	r := ToBase64([]byte("value"))
	fmt.Println(r.A)
	// Output:
	// dmFsdWU=
}

func ExampleToBase64String() {
	r := ToBase64("value")
	fmt.Println(r.A)
	fmt.Println(r.IsNil)
	// Output:
	// dmFsdWU=
	// false
}

func ExampleToBase64Nil() {
	r := ToBase64(nil)
	fmt.Println(r.IsNil)
	// Output:
	// true
}

func ExampleBase64To() {
	r := Base64To("dmFsdWU=")
	r2 := ToString(r.A)
	fmt.Println(r.Error)
	fmt.Println(r2.A)
	// Output:
	// <nil>
	// value
}

func ExampleBase64ToInt() {
	r := Base64To(100)
	r2 := ToString(r.A)
	fmt.Println(r.Error)
	fmt.Println(r2.A)
	// Output:
	// <nil>
}

func ExampleBase64ToNil() {
	r := Base64To(nil)
	r2 := ToString(r.A)
	fmt.Println(r.Error)
	fmt.Println(r.IsNil)
	fmt.Println(r2.A)
	// Output:
	// <nil>
	// true
}
