package main

import "fmt"

type info struct {
	Addss string
	Phone string
}

type user struct {
	ID   int
	Name string
	info
}

type employee struct {
	ID   int
	Name string
	info
}

func main() {
	u := user{
		ID:   1,
		Name: "user1",
		info: info{
			Addss: "addss_user",
			Phone: "phone_user",
		},
	}
	e := employee{
		ID:   1,
		Name: "user1",
		info: info{
			Addss: "addss_employee",
			Phone: "phone_employee",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
