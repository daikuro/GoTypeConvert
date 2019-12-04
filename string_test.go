package typeconv

import (
	"fmt"
	"io"
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

func ExampleArrayString() {
	bs := []byte("value")
	s := ToString(bs)
	fmt.Println(bs)
	fmt.Println(s.A)
	// Output:
	// [118 97 108 117 101]
	// value
}

func ExampleIoReaderNil() {
	var r io.Reader
	s := ToString(r)
	fmt.Println(s.A)
	fmt.Println(s.Error)
	// Output:
	// <nil>
}
