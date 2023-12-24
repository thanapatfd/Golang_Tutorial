package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", " C++"}
	fmt.Println(courseName)
	courseName = append(courseName, " Python ", "HTML", " C", " JS")
	fmt.Println(courseName)

	courseWeb := courseName[3:5]
	fmt.Println(courseWeb)

	courseWeb = courseName[:3]
	fmt.Println(courseWeb)

}
