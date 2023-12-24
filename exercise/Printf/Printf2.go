package main

import "fmt"

func main() {
	title := "Avengers: Endgame"
	year := 2019
	ratings := 8.4
	genre := "Sci-Fi"
	isSuperHero := true
	fav := '⭐'

	fmt.Printf("เรื่อง: %s\n", title)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรตติ้ง: %1.f\n", ratings)
	fmt.Printf("ประเภท: %s \n", genre)
	fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuperHero)
	fmt.Printf("เรื่องที่ชอบ: %c\n", fav)

}

// %s string , %q string เหมือนกันแต่มี "" คลุม
// %v - แสดงค่า default format ของ type นั้นๆ
