package typeconv

import (
	"fmt"
	"time"
)

func ExampleToTime() {
	t := ToTime(time.Date(2019, 1, 2, 3, 4, 5, 0, time.UTC))
	fmt.Println(t.A)
	// Output:
	// 2019-01-02 03:04:05 +0000 UTC
}

func ExampleToTimeDefault() {
	t := ToTime(nil, time.Date(2019, 1, 2, 3, 4, 5, 0, time.UTC))
	fmt.Println(t.A)
	// Output:
	// 2019-01-02 03:04:05 +0000 UTC
}
