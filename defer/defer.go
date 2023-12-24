package main

import "fmt"

func add(value1, value2 int) {
	result := value1 + value2
	fmt.Println("result :", result)

}
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}
}
func derferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j:", j)

	}
}
func main() {
	// fmt.Println("welcome to calculator")
	// defer fmt.Println("End")
	// defer add(20, 10)
	// defer add(15, 15)
	// defer add(12, 12)
	// defer add(2, 2)
	loop()
	derferloop()

}

// คำสั่ง defer เอาไว้ท้ายสุดของการทำงาน LFIO Last in ,First out สั่วนไหนที่สั่งให้ทำงานหลังสุด ก็จะไปทำงานส่วนแรก
