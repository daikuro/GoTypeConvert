package typeconv

import (
	"fmt"
	"time"
)

func ExampleMapToInterface() {
	type User struct {
		Name    string
		RegTime time.Time
	}
	v := MapToInterface(&User{}, map[string]interface{}{
		"Name":    "testUser",
		"RegTime": time.Now(),
	})
	fmt.Println(v.(*User).Name)
	// Output:
	// testUser
}
