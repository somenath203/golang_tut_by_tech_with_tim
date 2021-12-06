package main

import "fmt"

func main() {
	fmt.Println("Welcome to my Quiz Game")

	// creating a variable of type string
	var name string = "Somenath"

	// creating a variable of type number(can be both (+) and (-))
	var num int = 2000

	// creating a variable of unsigned(only positive) integer
	var unum uint = 7

	// creating a variable of type float (can be both (+) and (-))
	var dnum float64 = 4.56

	// creating a variable of type boolean
	var truefalse bool = true

	name = "Som" // changing the data that is stored inside 'name' variable.

	// displaying all the defined datas.
	fmt.Println("Name:-", name)
	fmt.Println("Number:-", num)
	fmt.Println("Unsigned Integer:-", unum)
	fmt.Println("Float Number:-", dnum)
	fmt.Println("Boolean Value:-", truefalse)
}
