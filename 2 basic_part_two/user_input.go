package main

import "fmt"

func main() {

	var name string
	var age int

	fmt.Print("Enter your Name:- ")
	fmt.Scan(&name)
	// Scan() is used to take user-input and the meaning of &name is that we want to store whatever the user has typed inside
	// variable name Here, &name actually refering to the memory location of variable name where the entered data needs to be stored.
	// Scan() requires the memory location where the value needs to be stored, therefore, we passed &name inside Scan().

	fmt.Print("Enter Age- ")
	fmt.Scan(&age)

	fmt.Println() // this will act like "\n".

	fmt.Println("Your Name:- ", name)
	fmt.Printf("Your Age:- %v ", age)

	//Scan() can only accept onw word and its execution stops when it receives '\n' after the entered string.

}
