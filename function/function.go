package main

import "fmt"

func hello() {
	fmt.Println("Hello world")
}

func main() {
	hello()
	result := plus(4, 3)
	fmt.Println("result :", result)

	result2 := plus3value(3, 4, 5)
	fmt.Println("result2 :", result2)
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}
func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}
