package main

import "fmt"

func main () {

	// arr := []int{1,2,3,4,5}


	// for i:= 0; i < len(arr); i++ {
	// 	fmt.Print(arr[i]," ")
	// }

	// using range
	// sum := 0

	// for index,num := range arr {
	// 	sum += num
	// 	fmt.Println("Index is ",index, " Number is ",num)

	// }
	// fmt.Println(sum)

	// for map
	// m := map[int]string{1:"Sahil",2:"Kumar"}

	// for k,v := range m {
	// 	fmt.Println("Key ",k, " Value ",v)
	// }

	// for string
	// unicode, code rune
	// starting byte of rune
	// in acii it is 255 byte
	// but in unicode it is more that acii value, if we put the value more that 255 in the first index then the index will shift Eg. if it exceedes 255 then it will not fit then instead of getting 0 -> 1, it will get 0 -> 2

	for i,v := range "golang" {
		fmt.Println(i,string(v))
	}
}