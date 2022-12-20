package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	current_user := User{"DJ", 12, false, "deepan@gmail.com"}
	fmt.Printf("The Current User is %v\n", current_user)
	fmt.Printf("The Current User is %+v\n", current_user)
	fmt.Println(current_user.GetStatus())
	fmt.Println(current_user.setEmail())
	fmt.Printf("The Current User is %+v\n", current_user)
}

type User struct {
	name   string
	age    int
	status bool
	email  string
}

func (u User) GetStatus() bool {
	fmt.Println("Is User Active ", u.status)
	return u.status
}

func (u User) setEmail() string {
	u.email = "test@gmail.com"
	fmt.Println("The Email is now set to ", u.email)
	return u.email
}
