package main

import "fmt"

func linePrinter(message string) {
	fmt.Println(message)
}

func linePrinter2(message string) {
	fmt.Println("Not your message!")
}

func picker(val int) func(string) {
	switch val {
	case 0:
		return linePrinter
	default:
		return linePrinter2
	}
}

func main() {
	var message = "This is a string"
	var printer1 = picker(0)
	var printer2 = picker(1)
	printer1(message)
	printer2(message)
}
