package typeconv

import (
	"fmt"
)

func ExampleToMap() {
	d := map[string]interface{}{}
	a := ToMap(d)
	fmt.Println(a)
	// Output:
	// map[]
}
