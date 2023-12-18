package main

import (
	"fmt"
	"time"
)

// func add(num1 int, num2 int) int {

// 	return num1 + num2
// }

// func subtract(num1, num2 int) int {
// 	result := num1 - num2

// 	if result < 0 {
// 		fmt.Println("Negative value")
// 	} else {
// 		fmt.Println("Positve Value")
// 	}
// 	return result
// }

// func named(name string) string {
// 	return name

// }

// type schemas struct {
// 	name string
// }

func displayHeader() string {
	message := fmt.Sprintf("AMOUNT  CATEGORY  DATE        DESCRIPTION")

	return message
}

func display(amount int, category string, date string, description string) string {
	message := fmt.Sprintf("N%d   %s  %s   %s", amount, category, date, description)

	return message
}

// var myArray []int

// // Initialize the array with values
// func init() {

// 	myArray = []int{6, 9, 3, 10, 7, 2, 7}

// }

// func mean(array []int) float64 {
// 	var addition int

// 	for i := 0; i < len(array); i++ {
// 		value := array[i]

// 		addition = addition + value

// 	}
// 	// Convert to float64 before division
// 	return float64(addition) / float64(len(array))
// }

func main() {
	fmt.Println("Welcome to Expense Tracker CLI V0.0.1")
	fmt.Println()

	fmt.Println("What is the amount")

	var amount int

	fmt.Scan(&amount)

	fmt.Println("What category is the expense")

	var category string

	fmt.Scan(&category)

	fmt.Println("Expense Description")

	var description string

	fmt.Scan(&description)

	var date = time.Now()

	// Format the date in a human-readable way
	currentDate := date.Format("2006-01-02")

	fmt.Println(displayHeader())

    fmt.Println()

	fmt.Println(display(amount, category, currentDate, description))

	// clock := time.Now()
	// fmt.Println("Hello World", clock)

}
