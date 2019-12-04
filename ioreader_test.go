package typeconv

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ExampleToIoReader() {
	r := ToIoReader("value")
	d, err := ioutil.ReadAll(r.A)
	if err != nil {
		panic(err)
	}
	fmt.Println(ToString(d).A)
	// Output:
	// value
}

func ExampleToIoReaderToString() {
	r := ToIoReader("value")
	fmt.Println(ToString(r.A).A)
	// Output:
	// value
}

func ExampleArrayByteToIoReaderToString() {
	r := ToIoReader([]byte("value"))
	fmt.Println(ToString(r.A).A)
	// Output:
	// value
}

func ExampleToIoReaderNil() {
	r := ToIoReader(nil)
	fmt.Println(r.A)
	// Output:
	// <nil>
}

func ExampleToIoReaderFile() {
	var f *os.File
	r := ToIoReader(f)
	fmt.Println(r.A)
	// Output:
	// <nil>
}
