package main

import "fmt"

type testStruct struct {
	firstName string
	lastName  string
}

func main() {
	var variable string
	variable = "Aaron"
	var pointer *string
	pointer = &variable
	fmt.Println(*pointer)

	var structure = testStruct{lastName: "Cunningham"}
	structure.firstName = variable
	fmt.Println(structure.firstName, structure.lastName)
}
