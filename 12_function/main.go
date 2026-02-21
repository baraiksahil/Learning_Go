package main

// func add(a int, b int) int {
// 	return a + b
// }

// func add(a, b int) int {
// 	return a + b
// }

// we can return multiple
// func languages()(string,string,string) {
// 	return "Js","Go","Java"
// }

// first class function

// we are reciving a fuc, inside the para we are reciving a int type and returning int
// func processIt(fn func(a int) int) {
// 	fn(1)
// }

// returning a func

func processIt() func(a int) int {
	return func(a int) int{
		return 4
	}
} 



func main() {
	// fmt.Println(add(10,20))
	// lang1, lang2, lang3 := languages()
	// lang1, lang2, _ := languages()
	// fmt.Println(lang1,lang2)

	// fn := func(a int)int  {
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt()

	fn(10)
}