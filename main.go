package main

import (
	"database/sql"
	"fmt"
	"github.com/mattn/go-sqlite3"
	// "time"
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

// func displayHeader() string {
// 	message := fmt.Sprintf("AMOUNT  CATEGORY  DATE        DESCRIPTION")

// 	return message
// }

// func display(amount int, category string, date string, description string) string {
// 	message := fmt.Sprintf("N%d   %s  %s   %s", amount, category, date, description)

// 	return message
// }

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
	database, _ := sql.Open("sqlite3", "./tracker.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXIST expense (id INTEGER PRIMARY KEY, amount INT, category TEXT, description TEXT, currentDate DATE)")

	statement.Exec()

	statement.Prepare("INSERT INTO expense (amount, category, description, currentDate) VALUES(?,?)")
	statement.Exec(100, "testCategory", "testDescription", "2023-12-21")
	rows, _ := database.Query("SELECT id, amount, category, description, currentDate FROM expense")

	var id int
	var amount int
	var category string
	var description string
	var currentDate Date

	for rows.Next() {
		rows.scan(&id, &amount, &category, &description, &currentDate)
		fmt.Println(strconv.Itoa(id) + ": " + amount + " " + category + " " + description + " " + currentDate)
	}

	// fmt.Println("Welcome to Expense Tracker CLI V0.0.1")
	// fmt.Println()

	// fmt.Println("What is the amount")

	// var amount int

	// fmt.Scan(&amount)

	// fmt.Println("What category is the expense")

	// var category string

	// fmt.Scan(&category)

	// fmt.Println("Expense Description")

	// var description string

	// fmt.Scan(&description)

	// var date = time.Now()

	// // Format the date in a human-readable way
	// currentDate := date.Format("2006-01-02")

	// fmt.Println(displayHeader())

	// fmt.Println()

	// fmt.Println(display(amount, category, currentDate, description))

	// clock := time.Now()
	// fmt.Println("Hello World", clock)

}
