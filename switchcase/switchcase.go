package main

import "fmt"

func main() {
	// input := 4
	// switch input {
	// case 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// case 3:
	// 	fmt.Println("3")
	// default:
	// 	fmt.Println("Default")
	// }
	var input string
	fmt.Print("Enter Color or -1 to Exit: ")
	fmt.Scan(&input)

	switch input {
	case "blue":
		fmt.Println("Color rgb is #0000FF")
	case "green":
		fmt.Println("Color rgb is #008000")
	case "pink":
		fmt.Println("Color rgb is #FFC0CB")
	case "yellow":
		fmt.Println("Color rgb is #FFFF00")
	default:
		fmt.Println("No Color")
	}

}
