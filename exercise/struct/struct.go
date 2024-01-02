package main

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

func main() {
	// Workshop: struct
	// กำหนด: 1. ให้นิยามโครงสร้างข้อมูล movie เพื่อเก็บ ชื่อเรื่อง(string) ปี(ตัวเลข) เรตติ้ง(ตัวเลขทศนิยม) ประเภท(slice ของ string) และ isSuperHero(bool)
	// 	  2. ให้ประกาศตัวแปรเพื่อเก็บหนัง Avengers: Endgame, ปี:2019, เรตติ้ง:8.4, ประเภท:["Action", "Drama"] และ isSuperHero:true
	//	  3. ให้ประกาศตัวแปรเพื่อเก็บหนัง Infinity War, ปี:2018, เรตติ้ง:8.4, ประเภท:["Action", "Sci-Fi"] และ isSuperHero:true
	//	  4. ให้เก็บหนังทั้งสองเรื่องไว้ในตัวแปร mvs
	// 	  5. ทำการวนลูปเพื่อแสดงผลหนังทีละเรื่อง
	//
	// Output:
	// main.movie{title:"Avengers: Endgame", year:2019, rating:8.4, genres:[]string{"Action", "Drama"}, isSuperHero:true}
	// main.movie{title:"Avengers: Infinity War", year:2018, rating:8.4, genres:[]string{"Action", "Sci-Fi"}, isSuperHero:true}

	// TODO: write code below.

	mvs1 := movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}
	mvs2 := movie{
		title:       "Avengers: Infinity War",
		year:        2018,
		rating:      8.4,
		genres:      []string{"Action", "Sci-Fi"},
		isSuperHero: true,
	}

	var mvs []movie
	mvs = []movie{mvs1, mvs2}

	for _, mv := range mvs {
		fmt.Println("%#v\n", mv)
	}

	// movielist = append(movielist, mvs1)

	// fmt.Println("title: ", mvs1.title)
	// fmt.Println("year: ", mvs1.year)
	// fmt.Println("rating: ", mvs1.rating)
	// fmt.Println("genres: ", mvs1.genres)
	// fmt.Println("isSuperhero: ", mvs1.isSuperHero)

}
