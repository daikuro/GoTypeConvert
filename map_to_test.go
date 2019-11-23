package typeconv

import "fmt"

func ExampleMapToInterface() {
	type User struct {
		Name string
	}
	v := MapToInterface(&User{}, map[string]interface{}{
		"Name": "testUser",
	})
	fmt.Println(v.(*User).Name)
	// Output:
	// testUser
}
