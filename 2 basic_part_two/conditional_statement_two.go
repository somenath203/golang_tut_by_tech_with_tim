package main

import "fmt"

func main() {

	var firstname string
	var lastname string

	var age uint

	score := 0
	marks_in_total := 5

	fmt.Println("<-----Coding Quiz----->")

	fmt.Println()

	fmt.Print("Enter Your Name [Only Two Words Are Accepted]:- ")
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
		fmt.Printf("Your current age is %v. You are NOT ELIGIBLE for the quiz.", age)
		return
		// here, wwe wrote 'return' inside else block. The work of 'return' is to completely break the program and come out of it.
		// It means if we write something after this 'return', both inside and outside of 'else' block, then, it won't execute.
	}

	fmt.Println()
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
		score++
	} else if ansone == "B" {
		fmt.Println("Your answer is Correct...")
		score++
	} else {
		fmt.Println("Your answer is Wrong...")
	}

	fmt.Println()

	fmt.Println("Q2) React is managed by which Company ?")
	fmt.Println("    a. Google      b.Microsoft")
	fmt.Println("    c. Nvidia      d.Facebook")

	var anstwo string

	fmt.Scan(&anstwo)

	if anstwo == "d" {
		fmt.Println("Your answer is Correct...")
		score++
	} else if anstwo == "D" {
		fmt.Println("Your answer is Correct...")
		score++
	} else {
		fmt.Println("Your answer is Wrong...")
	}

	fmt.Println()

	fmt.Println("Q3) Which one is a backend framework made with Python?")
	fmt.Println("    a. PHP          b.Flask")
	fmt.Println("    c. Express      d.Vue")

	var ansthree string

	fmt.Scan(&ansthree)

	if ansthree == "b" {
		fmt.Println("Your answer is Correct...")
		score++
	} else if ansthree == "B" {
		fmt.Println("Your answer is Correct...")
		score++
	} else {
		fmt.Println("Your answer is Wrong...")
	}

	fmt.Println()

	fmt.Println("Q4) Which framework have been ranked as 'The Most Loved Framework' for creating single page website by Stack Overflow Developer Survey 2021 ?")
	fmt.Println("    a. Svelte      b.React")
	fmt.Println("    c. Angular      d.Vue")

	var ansfour string

	fmt.Scan(&ansfour)

	if ansfour == "a" {
		fmt.Println("Your answer is Correct...")
		score++
	} else if ansfour == "A" {
		fmt.Println("Your answer is Correct...")
		score++
	} else {
		fmt.Println("Your answer is Wrong...")
	}

	fmt.Println()

	fmt.Println("Q5) ES5 came out in which year ?")
	fmt.Println("    a. 2015      b.2016")
	fmt.Println("    c. 2009      d.2013")

	var ansfive string

	fmt.Scan(&ansfive)

	if ansfive == "c" {
		fmt.Println("Your answer is Correct...")
		score++
	} else if ansfive == "C" {
		fmt.Println("Your answer is Correct...")
		score++
	} else {
		fmt.Println("Your answer is Wrong...")
	}

	fmt.Println()

	fmt.Println("<-----RESULT----->")
	fmt.Printf("Your Score :- %v/%v",score, marks_in_total)

	fmt.Println()

	percent := (float64(score) / 5.0) * 100 // here, '5' is no.of questions.
	fmt.Printf("Your Percentage :- %v%%", percent)
	
}
