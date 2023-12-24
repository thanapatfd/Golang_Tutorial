package main

import "fmt"

var productName [4]string

func main() {
	productName[0] = "Macbook"
	productName[1] = "Iphone"
	productName[2] = "Ipad"
	productName[3] = "Airpods"

	price := [4]float32{40000, 30000, 0, 0}

	fmt.Println(productName)
	fmt.Println(price)

}
