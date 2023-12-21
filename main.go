package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// "time"

	_ "github.com/mattn/go-sqlite3"
)

// type schemas struct {
// 	name string
// }

// Format and Display Expenses
func display(id int, amount int, category string, date string, description string) string {
	message := fmt.Sprintf("%d    N%d   %s  %s   %s", id, amount, category, date, description)

	return message
}

// Add Expenses
func addExpense(database *sql.DB, err error) {
	fmt.Println("What's the amount")

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

	var statement, _ = database.Prepare("INSERT INTO expense (amount, category, description, currentDate) VALUES(?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(amount, category, description, currentDate)
	fmt.Println("Expense added succesfully")
}

// Get All Expenses
func getAllExpenses(rows *sql.Rows, id int, amount int, category string, description string, currentDate string) {
	fmt.Println("ID   AMOUNT CATEGORY       DATE         DESCRIPTION")

	fmt.Println()

	rows.Next()
	{
		rows.Scan(&id, &amount, &category, &description, &currentDate)

		fmt.Println(display(id, amount, category, currentDate, description))
	}
}

func main() {
	database, _ := sql.Open("sqlite3", "./tracker.db")

	database.SetMaxOpenConns(1)

	var statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS expense (id INTEGER PRIMARY KEY, amount INTEGER, category TEXT, description TEXT, currentDate TEXT)")

	statement.Exec()

	rows, _ := database.Query("SELECT id, amount, category, description, currentDate FROM expense")

	var id int
	var amount int
	var category string
	var description string
	var currentDate string

	var command, verb string
	// defer statement.Exec(amount, category, description, &currentDate)

	fmt.Println("Welcome to Expense Tracker CLI V0.0.1")
	fmt.Println()

	fmt.Println("add the option -help for instructions e.g tracker -help")

	fmt.Println()

	fmt.Scan(&command, &verb)

	if command == "tracker" && verb == "-help" {
		fmt.Println()
		fmt.Println("option -ls for listing all the expenses")
		fmt.Println()
		fmt.Println("option -add to add new expenses")
		fmt.Println()
		fmt.Println("option -clear to delete all records in expense db")
		fmt.Println()
	} else if command == "tracker" && verb == "-ls" {
		getAllExpenses(rows, id, amount, category, description, currentDate)

	} else if command == "tracker" && verb == "-add" {
		addExpense(database, err)
	} else {
		fmt.Println("Command not found. try 'tracker -help' for instructions")
	}

}
