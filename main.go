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

func main() {
	clock := time.Now()
	fmt.Println("Hello World", clock)
	fmt.Println(add(5, 4))
	fmt.Println(subtract(2, 4))
	fmt.Println(named("Emmanuel"))

	person1 := schemas{name: "PROGRAMMED"}

	fmt.Println(person1.display())
}
