package main

import "fmt"

func main() {
	rate := 8.4
	switch {
	case rate >= 7.0 && rate < 10:
		fmt.Println("Good 🥰")
	case rate >= 5.0 && rate < 7.0:
		fmt.Println("Normal 😐")
	case rate < 5.0:
		fmt.Println("Disappointed 😞")
	default:
		fmt.Println("🤷🤷🤷🤷")

	}
}

// if rate >= 7.0 && rate < 10.0 {
// 	fmt.Println("Good 🥰")
// } else if rate >= 5.0 && rate < 7.0 {
// 	fmt.Println("Normal 😐")
// } else if rate < 5.0 {
// 	fmt.Println("Disappointed 😞")
// } else {
// 	fmt.Println("🤷🤷🤷🤷")
// }
