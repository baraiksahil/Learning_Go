package main

import "fmt"


func sum(nums ...int)int {
	total := 0

	for _,v := range nums {
		total += v
	}
	return total
}

func main() {
	// result := sum(10,30,50)

	nums := []int{10,20,30}
	result := sum(nums...)
	fmt.Println(result)

}