package main


func main(){

	// normal switch
	// i := 10
	// switch i {
	// case 1:
	// 	println("One")
	// case 2:
	// 	println("Two")
	// default:
	// 	println("Others")
	// }

	// multipe switch cindition
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("Its weekend")
	// default:
	// 	println("Its work day")
	// }

	// type switch
	whoAmI := func(dataType interface{}){
		switch ty := dataType.(type) {
		case int:
			println("Integer",ty)
		case bool:
			println("Boolean",ty)
		case float64:
			println("Float",ty)
		default:
			println("Others",ty)
		}
	}

	whoAmI(10.0)
}