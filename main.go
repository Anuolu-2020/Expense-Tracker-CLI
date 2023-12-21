package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	// "time"

	_ "github.com/mattn/go-sqlite3"
)

// type schemas struct {
// 	name string
// }

// Format and Display Expenses
func display(id int, amount int, category string, description string, date string) string {
	message := fmt.Sprintf("%d    N%d   %s  %s   %s", id, amount, category, description, date)

	return message
}

// Add Expenses
func addExpense(database *sql.DB) {
	fmt.Println("What's the amount")

	var amount int

	fmt.Scan(&amount)

	fmt.Println("What category is the expense")

	var category string

	fmt.Scan(&category)

	fmt.Println("Expense Description")

	var description string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		description = scanner.Text()
	}

	var date = time.Now()

	// Format the date in a human-readable way
	currentDate := date.Format("2006-01-02")

	statement, err := database.Prepare("INSERT INTO expense (amount, category, description, currentDate) VALUES (?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	statement.Exec(amount, category, description, currentDate)

	fmt.Println()

	fmt.Println("Expense added succesfully")
}

func getAllExpenses(rows *sql.Rows) {
	fmt.Println("ID   AMOUNT  CATEGORY  DESCRIPTION        DATE")
	fmt.Println()

	for rows.Next() {
		var id int
		var amount int
		var category string
		var description string
		var currentDate string

		err := rows.Scan(&id, &amount, &category, &description, &currentDate)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(display(id, amount, category, description, currentDate))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	database, err := sql.Open("sqlite3", "./tracker.db")

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	var statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS expense (id INTEGER PRIMARY KEY, amount INTEGER, category TEXT, description TEXT, currentDate TEXT)")

	statement.Exec()

	rows, _ := database.Query("SELECT id, amount, category, description, currentDate FROM expense")

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
		getAllExpenses(rows)

	} else if command == "tracker" && verb == "-add" {
		addExpense(database)
	} else if command == "tracker" && verb == "-clear" {
		statement, err := database.Prepare("DELETE FROM expense")

		if err != nil {
			log.Fatal(err)
		}

		defer statement.Close()

		_, err = statement.Exec()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Expenses data deleted")
	} else {
		fmt.Println("Command not found. try 'tracker -help' for instructions")
	}

}
