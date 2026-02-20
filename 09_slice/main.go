package main

import (
	"fmt"
	"slices"
)

// dynamic size
// most used concept in go
// + usefull methods
func main() {
	// // uninitialize slice == nil
	// var slice []int;

	// fmt.Println(slice == nil) // true

	// // another way how to make a slice
	// slice = make([]int, 0,5) // initial size, capacity

	// // append in slice
	// slice = append(slice, 1)
	// slice = append(slice, 2)
	// slice = append(slice, 3)
	// slice = append(slice, 4)
	// slice = append(slice, 5)
	// slice = append(slice, 6)

	// fmt.Println(slice)
	// fmt.Println(cap(slice)) // capacity of slice
	// fmt.Println(len(slice)) // len of slice

	// // copy a slice
	// var nums = make([]int,len(slice))

	// // detination, src
	// copy(nums,slice)
	// fmt.Println(nums,slice)

	var slice = []int{1,2,3,4}
	var nums = []int{1,2,3,4,5}

	// slice operator
	// it will print value from 0 -> 2, it will exclude 3
	// fmt.Println(slice[0:])

	// slices packege
	fmt.Println(slices.Equal(slice,nums))

	// 2D slices
	var twoDslice = [][]int{{1,2,3},{1,2}}

	fmt.Println(twoDslice)

}