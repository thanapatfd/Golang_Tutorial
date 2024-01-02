package main

import "fmt"

func main() {

	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}

	xs = append(xs, ys...)

	vote := []float64{}

	vote = append(vote, xs...)

	fmt.Println("vote: ", vote)

	voteSlice := vote[5:9]

	fmt.Println("vote 5 - 8: ", voteSlice)
}
