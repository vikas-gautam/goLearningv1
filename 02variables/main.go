package main
import "fmt"

//variables, types and constant

const LoginToken string = "xyzxyz" // here first letter Caps means equals to public const loginToken

func main() {
	fmt.Println("variables")
	var username string = "vikas"
	fmt.Println(username)
	fmt.Printf("the variable is of type: %T \n", username)

	// implicit type 

	var anothervariable = "somevalue"    // here we have not defined variable type so lexer will do that for u
	fmt.Println(anothervariable)

	//no var style but only allowed inside the method

	numberOfNinjas := 20
    fmt.Println(numberOfNinjas)

	fmt.Println(LoginToken)

}