package main

import "fmt"

func main() {
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}

func emote(r float64) string {
	switch {
	case r < 5.0:
		return "Disappointed 😞"
	case r >= 5.0 && r < 7.0:
		return "Normal 😐"
	case r >= 7.0 && r < 10.0:
		return "Good 🥰"
	default:
		return "🤷🤷🤷🤷"
	}
}
