package main

import (
	"fmt"
	"time"
)

func add(num1 int, num2 int) int {

	return num1 + num2
}

func subtract(num1, num2 int) int {
	result := num1 - num2

	if result < 0 {
		fmt.Println("Negative value")
	} else {
		fmt.Println("Positve Value")
	}
	return result
}

func named(name string) string {
	return name

}

type schemas struct {
	name string
}

func (data schemas) display() string {
	return data.name
}

var myArray []int

// Initialize the array with values
func init() {

	myArray = []int{6, 9, 3, 10, 7, 2, 7}

}

func mean(array []int) float64 {
	var addition = 0

	for i := 0; i < len(array); i++ {
		value := array[i]

		addition = addition + value

	}
	// Convert to float64 before division
	return float64(addition) / float64(len(array))
}

func main() {
	clock := time.Now()
	fmt.Println("Hello World", clock)
	fmt.Println(add(5, 4))
	fmt.Println(subtract(2, 4))
	fmt.Println(named("Emmanuel"))
	fmt.Println(mean(myArray[:]))

	person1 := schemas{name: "PROGRAMMED"}

	fmt.Println(person1.display())
}
