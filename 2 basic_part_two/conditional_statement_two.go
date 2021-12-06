package main

import "fmt"

func main() {

	var firstname string
	var lastname string

	var age uint

	fmt.Println("<-----Coding Quiz----->")

	fmt.Println()

	fmt.Print("Enter Your Full Name:- ")
	fmt.Scan(&firstname, &lastname)

	fmt.Println()

	fmt.Printf("Hello %v %v, welcome to the game. \n", firstname, lastname)

	fmt.Println()

	fmt.Print("Enter Your Age:- ")
	fmt.Scan(&age)

	fmt.Println()

	if age >= 10 {
		fmt.Printf("Your current age is %v. You are ELIGIBLE for the quiz.", age)
	} else {
		fmt.Printf("Your current age is %v. You are NOT ELIGIBLE for the quiz. \n", age)
		return
		// here, wwe wrote 'return' inside else block. The work of 'return' is to completely break the program and come out of it.
		// It means if we write something after this 'return', both inside and outside of 'else' block, then, it won't execute.
	}

	fmt.Println()

	fmt.Println("Here We GO.....")

	fmt.Println()

	fmt.Println("Click on the Right Option")

	fmt.Println()

	fmt.Println("Q1) Who is the founder of C ?")
	fmt.Println("    a. Bjarne Stroustrup  b.Dennis Ritchie")
	fmt.Println("    c. Jordan Walke       d.None of the Above")

	var ansone string

	fmt.Scan(&ansone)

	if ansone == "b" {
		fmt.Println("Your answer is Correct...")
	} else {
		fmt.Println("Your answer is Wrong...")
	}
}
