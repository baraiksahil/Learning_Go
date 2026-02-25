package main

import "fmt"

// by value
// func changeNumber (num int) int {
// 	num = 5
// 	fmt.Println("In change number",num)
// 	return num
// }

// by reference
// receiving pointer
func changeNumber(num *int) {
	// dereferencing
	*num = 5
	// fmt.Println("In change number",num)
	// if we want to see value
	fmt.Println("In change number",*num)
}

func main() {
	num := 1

	// var ans = changeNumber(num)
	// fmt.Println("In main func",num)
	// how to print the address
	// fmt.Println("Memory address",&num)
	// fmt.Println("Answer is",ans)
	changeNumber(&num)
	fmt.Println(num)

}