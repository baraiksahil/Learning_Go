package main

import "fmt"

// We can declare datatype outside the main fun

const name = "Sahil"


// this is allowed
// var age int

// this is not allowed
// age := 10

func main() {

	// this is not allowed in constant
	// const name string
	// name = "Sahil"

	// const name string = "Sahil"

	
	// age = 10;

	const (
		PORT = 3000
		host = "localhost"
	)

	fmt.Println(PORT,host)
}