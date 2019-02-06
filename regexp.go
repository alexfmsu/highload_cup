package main

import (
	"encoding/json"
	"fmt"
)

func MARSHALL() {
	type Artist struct {
		Id int
	}

	temp := []byte(`{"Id":12}`)

	var artist Artist
	err := json.Unmarshal(temp, &artist)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Println(artist.Id)
	fmt.Printf("%+v", artist)
}

func main() {
	MARSHALL()
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // "regexp"

// type Account struct {
// 	id      uint32
// 	email   string
// 	fname   string
// 	sname   string
// 	phone   string
// 	sex     string
// 	birth   string
// 	country string
// 	city    string
// }

// type Message struct {
// 	Name string
// 	Body string
// 	Time int64
// }

// func main() {
// 	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
// 	var m Message

// 	print(b)

// 	if err := json.Unmarshal(b, &m); err != nil {
// 		fmt.Printf("%+v", m)
// 	}

// 	// re := regexp.MustCompile(`[^/]+/recommend`)

// 	// s := "1/recommend"

// 	// if re.MatchString(s) {
// 	// 	println("match")
// 	// } else {
// 	// 	println("not match")
// 	// }
// }
