package main

import (
	"fmt"
)

func main() {
	var (
		a, b     float64
		operator string
	)

	//Get input for variable 'a'
	fmt.Print("Enter A : ")
	fmt.Scan(&a)

	//Get input for variable 'b'
	fmt.Print("Enter B : ")
	fmt.Scan(&b)

	//Get input for operator
	fmt.Print("Enter operator .eg. +, -, x, / : ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		add(a, b)
	case "-":
		sub(a, b)
	case "x":
		mul(a, b)
	case "/":
		div(a, b)
	default:
		fmt.Println("Enter a valid option.")
	}

}

func add(x, y float64) {
	printAnwser(x+y, "+")
}

func sub(x, y float64) {
	printAnwser(x-y, "-")
}

func mul(x, y float64) {
	printAnwser(x*y, "x")
}

func div(x, y float64) {
	printAnwser(x/y, "/")
}

func printAnwser(value float64, operator string) {
	fmt.Printf("A %v B = %v\n", operator, value)
}
