package main

import "fmt"

func main() {
	rate := 8.4
	if rate >= 7.0 && rate < 10. {
		fmt.Println("Good 🥰")
	} else if rate >= 5.0 && rate < 7.0 {
		fmt.Println("Normal 😐")
	} else if rate < 5.0 {
		fmt.Println("Disappointed 😞")
	} else {
		fmt.Println("🤷🤷🤷🤷")
	}
}
