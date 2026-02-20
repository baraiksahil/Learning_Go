package main

import "fmt"
func main() {
	var nums [4]int

	// - when to use array = fixed size
	// - memory optimazation
	// - constant time access
	
	// array lenght
	fmt.Println(len(nums))

	// how to add element
	nums[0] = 1
	fmt.Println(nums[0])

	// how to print whole array
	fmt.Println(nums)

	var isTrue [4]bool 
	// it will print falsey values
	fmt.Println(isTrue)

	var name [4]string
	// it will print empty values
	fmt.Println(name) 

	// to declare array inn single line
	arr := [3]int{10,20,30}
	fmt.Println(arr)

	// 2D array

	twoDarr := [3][3]int{{10,20,30},{10,30,40}}
	fmt.Println(twoDarr)
}