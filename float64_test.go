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
