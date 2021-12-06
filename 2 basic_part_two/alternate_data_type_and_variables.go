package main

import "fmt"

func main() {
	fmt.Println("Welcome to my Quiz Game")

	name := "Som" //':= ' implicitly detects the datatype of the variable stored inside the variable 'name' for which we don't have to
	              //      define variable all the time.
				  
    age := 21
	
	fmt.Printf("Hello %v %v, your current age is %v.", name, name, age) 
	// to display anything using "%v" or any placeholder, we use 'Printf' instead of 'Println'.
	// "%v" is going to take the value about what the variable 'name' has in this case.
	// %d stands for decimal notation.

	// placeholder/formatting documentation :- https://gobyexample.com/string-formatting

}
