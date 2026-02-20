package main

import (
	"fmt"
	"maps"
)



func main() {

	// how to declare + initialize the variable in map
	// var name = map[int]string{
	// 	1:"Sahil",
	// 	2:"Raj",
	// 	3:"Kunal",
	// }

	// fmt.Println(name)

	// how to delete element in map
	// delete(name,2)

	// how to clear all the elements
	// clear(name)

	// fmt.Println(name)

	// how to make map using make

	// mpp := make(map[int]string)

	// mpp[1] = "golang"
	// mpp[2] = "backend"

	// fmt.Println(mpp)

	// how to get element
	// if we want to access the key which is not present if string -> empty, int -> 0, bool -> false

	// fmt.Println(mpp[1])

	// how to check for the element
	var mpp = map[string]int{"age":26,"phone":12345565}
	fmt.Println(mpp)

	value,ok := mpp["age"]

	if ok {
		fmt.Println("all ok",value)
	} else {
		fmt.Println("not ok",value) // if not ok then it will show 0,empty,false
	}

	m1 := map[int]string{1:"Sahil",2:"Kumar"}
	m2 := map[int]string{1:"Sahil",3:"Kumar"}

	fmt.Println(maps.Equal(m1,m2))

}