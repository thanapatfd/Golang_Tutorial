package main

import "fmt"

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i = ", i)
	// 	i += 1
	// }

	// for j := 0; j < count; j++ {
	// 	fmt.Println("j = ", j)
	// }
	// count := 10
	// for count > 5 {
	// 	fmt.Println("Hello Test")
	// 	count--
	// }
	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("Input = ", input)
		if input == "-1" {
			break
		}
	}
}
