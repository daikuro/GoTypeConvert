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

func ExampleToMapInt() {
	d := map[string]interface{}{
		"A": 111,
	}
	a := ToMap(d)
	fmt.Println(a)
	// Output:
	// map[A:111]
}

func ExampleToMapStruct() {
	type User struct {
		Name string
	}
	user := &User{
		Name: "TestName",
	}

	a := ToMap(user)
	fmt.Println(a)
	// Output:
	// map[Name:TestName]
}

func ExampleToMapStructFields() {
	type User struct {
		Name string
		Id   int
	}
	user := &User{
		Name: "TestName",
		Id:   100,
	}

	a := ToMap(user)
	fmt.Println(a["Id"], a["Name"])
	// Output:
	// 100 TestName
}

func ExampleToMapStructNest() {
	type User struct {
		Name    string
		SubUser *User
	}
	user := &User{
		Name: "TestName",
		SubUser: &User{
			Name: "SubName",
		},
	}

	a := ToMap(user)
	//fmt.Println(a["SubUser"].(*User).Name)
	fmt.Println(a["Name"], a["SubUser"])
	// Output:
	// TestName map[Name:SubName]
}

func ExampleToMapStructPrivateField() {
	type User struct {
		Name string
		id   int
	}
	user := &User{
		Name: "TestUser",
		id:   100,
	}

	a := ToMap(user)
	fmt.Println(a)
	// Output:
	// map[Name:TestUser]
}

func ExampleToMapString() {
	a := ToMap("aaa")
	fmt.Println(a)
	// Output:
	// map[]
}
