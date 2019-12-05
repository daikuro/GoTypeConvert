package typeconv

import (
	"fmt"
	"time"
)

func ExampleToMap() {
	d := map[string]interface{}{}
	a := ToMap(d)
	fmt.Println(a.A)
	fmt.Println(a.IsNil)
	fmt.Println(a.Error)
	// Output:
	// map[]
	// false
	// <nil>
}

func ExampleToMapNil() {
	a := ToMap(nil)
	fmt.Println(a.A)
	fmt.Println(a.IsNil)
	// Output:
	// map[]
	// true
}

func ExampleToMapStruct() {
	type User struct {
		Name string
	}
	user := &User{
		Name: "TestName",
	}

	a := ToMap(user).A
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

	a := ToMap(user).A
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

	a := ToMap(user).A
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

	a := ToMap(user).A
	fmt.Println(a)
	// Output:
	// map[Name:TestUser]
}

func ExampleToMapStructNil() {
	type User struct {
		Name interface{}
	}
	user := &User{
		Name: nil,
	}
	a := ToMap(user)
	fmt.Println(a.A)
	fmt.Println(a.IsNil)
	fmt.Println(a.Error)
	// Output:
	// map[]
	// false
	// <nil>
}

func ExampleToMapString() {
	a := ToMap("aaa")
	fmt.Println(a.A)
	fmt.Println(a.NotStruct)
	// Output:
	// map[]
	// true
}

func ExampleToMapArray() {
	d := []string{"a", "b"}
	a := ToMap(d)
	fmt.Println(a.A)
	fmt.Println(a.NotStruct)
	// Output:
	// map[]
	// true
}

func ExampleToMapTime() {
	type User struct {
		OrderTime time.Time
	}
	user := &User{}
	a := ToMap(user)
	if a.Error != nil {
		fmt.Printf("err %+v", a.Error)
		//panic(a.Error)
	}
	fmt.Println(a.A)
	// Output:
	// map[OrderTime:0001-01-01 00:00:00 +0000 UTC]
}

func ExampleToMapTimePtrNil() {
	type User struct {
		OrderTime *time.Time
	}
	user := &User{}
	a := ToMap(user)
	if a.Error != nil {
		fmt.Printf("err %+v", a.Error)
		//panic(a.Error)
	}
	fmt.Println(a.A)
	// Output:
	// map[]
}

func ExampleToMapTimePtr() {
	type User struct {
		OrderTime *time.Time
	}
	t := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	user := &User{
		OrderTime: &t,
	}
	a := ToMap(user)
	if a.Error != nil {
		fmt.Printf("err %+v", a.Error)
		//panic(a.Error)
	}
	fmt.Println(a.A)
	// Output:
	// map[OrderTime:0000-01-01 00:00:00 +0000 UTC]
}
