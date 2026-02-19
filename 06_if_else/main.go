package main


func main() {

	// age := 18
	// if age >= 18 {
	// 	println("Person is an adult",age)
	// } else {
	// 	println("Person is not an adult",age)
	// }

	// age := 10
	// if age >= 18 {
	// 	println("Person is an adult",age)
	// } else if age >= 12 {
	// 	println("Person is a teenage",age)
	// } else {
	// 	println("Person is a kid",age)
	// }

	// in go we don't have ternary operator

	if age := 10; age >= 18 {
		println("Person is an adult",age)
	} else if age >= 12 {
		println("Person is a teenage",age)
	} else {
		println("Person is a kid",age)
	}
}