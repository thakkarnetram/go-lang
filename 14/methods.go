package main

import "fmt"

func main() {
	fmt.Println("Methods in GO ")
	devil:=User{"NewName","Email@sa",8,"Ojke"}
	devil.GetStatus()
	devil.NewEmail()
}

type User struct {
	Name string
	Email string
	Age int
	Password string
}
// method
func (user User) GetStatus () {
	fmt.Println("Does user have email " ,  user.Email)
}

// new method
func(user User) NewEmail(){
	user.Email = "newEmaiL@dev"
	fmt.Println("Email of user is changed to -> ", user.Email) // this creates a copy and doesnt change the original email
}