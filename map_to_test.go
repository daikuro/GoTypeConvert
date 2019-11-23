package typeconv

import (
	"fmt"
	"log"
	"time"
)

func ExampleMapToInterface() {
	type User struct {
		Name    string
		RegTime *time.Time
	}
	t := time.Now()
	v := MapToInterface(&User{}, map[string]interface{}{
		"Name":    "testUser",
		"RegTime": &t,
	})

	// check
	user := v.(*User)
	if user.RegTime != &t {
		log.Panic("error ", user.RegTime, t)
	}
	fmt.Println(user.Name)
	// Output:
	// testUser
}

func ExampleMapToInterfaceNilValue() {
	type User struct {
		Name    string
		RegTime *time.Time
	}
	v := MapToInterface(&User{}, map[string]interface{}{
		"Name":    "testUser",
		"RegTime": nil,
	})
	fmt.Println(v.(*User).RegTime)
	// Output:
	// <nil>
}
