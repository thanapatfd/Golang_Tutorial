package main

import "fmt"

func main() {
	genres := [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("before for loop: %#v\n", genres)
	// for loop here.
	for i := 0; i < len(genres); i++ {
		genres[i] = "Movie : " + genres[i]
	}
	fmt.Printf("after  for loop: %#v\n", genres)

}
