package main

import "fmt"

type testStruct struct {
	firstValue  int
	secondValue int
}

func main() {
	var structure = testStruct{firstValue: 4, secondValue: 6}
	var x, y, output, _ = adder(structure)
	fmt.Println(x, y, output)
}

func adder(input testStruct) (int, int, int, string) {
	return input.firstValue, input.secondValue, (input.firstValue + input.secondValue), "Fourth return"
}
