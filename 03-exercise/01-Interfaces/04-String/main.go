package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return "Name: " + u.name + " | Email: " + u.email
}

// TODO: Implement custom formating for user struct values.

func main() {
	u := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}
	fmt.Println(u)
}
