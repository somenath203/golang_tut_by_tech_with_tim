package main

import "fmt"

func main(){

	fmt.Println("<------Coding Quiz------>")
	fmt.Println()

	var name string
	var age uint // 'uint' means unsiged integer means we don't want anyone to enter negative age.

	fmt.Print("Enter your Name:- ")
	fmt.Scan(&name)

	fmt.Println()

	fmt.Print("Enter your Age:- ")
	fmt.Scan(&age)

	fmt.Println()

	fmt.Printf("Welcome to the Quiz %v.",name)

	fmt.Println()

	fmt.Println("Old Enought to play? ",age >= 10) 
	// Here, we are bascially checking wether 'age' is greaterthan or equals to 10 or not. If 'age' is greaterthan or equals to 10,
	// then, 'true' will be printed, else 'false'.
}

