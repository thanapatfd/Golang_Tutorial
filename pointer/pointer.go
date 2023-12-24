package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}
func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)
	zerovalue(i)
	fmt.Println("i from fuction zerovalue =", i)
	zeropointer(&i)
	fmt.Println("i value fuction zerovalue =", i)
	fmt.Println("i address fuction zerovalue =", &i)

}
