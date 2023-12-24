package main

import "os"

func main() {
	data1 := []byte("Hi\nborntoDev")
	err := os.WriteFile("data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("employeeName")

	check(err)
	defer f.Close()

	data2 := []byte("Sira\nThanapat")
	os.WriteFile("employeeName.txt", data2, 0644)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
