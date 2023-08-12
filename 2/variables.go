package main

import "fmt"

const LoginToken string = "sasaewasa"
const jwtToken string = "jwt"

// ! Starting with a uppercase letter goLang marks it as Public:

func main() {
	// ! Inside methods we can use :=  to assign stuff but not outside/ globally
	var name string = "new"
	fmt.Println(name)
	fmt.Printf("type is : %T \n" , name)

	const email string = "ok@mail.com"
	fmt.Println(email)
	fmt.Printf("type email is : %T \n" , email )

	var isAuth  = true
	fmt.Println(isAuth)
	fmt.Printf("type of isAuth : %T \n", isAuth)

	// default val and aliases 
	var defaultValue string
	fmt.Println(defaultValue)
	fmt.Printf("type of default  : %T \n", defaultValue)

	// implict
	var implict = "ok22"
	fmt.Println(implict)
	fmt.Printf("type of implict  : %T \n", implict)

	// no var style
	users := 33;
	fmt.Println("Users ",users)
	fmt.Printf("type of users  : %T \n", users)

	// global constants
	fmt.Println("Login token " , LoginToken)
	fmt.Println("JWT token " , jwtToken)
}