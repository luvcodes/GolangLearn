package main

import (
	"encoding/json"
	"fmt"
)

// Movie define with json tags
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"Casablanca", 1942, false, []string{"Humphrey Bogart", "Ingrid Bergman"}}

	// 编码的过程, 结构体 -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json.Marshal failed, err:", err)
		return
	}

	fmt.Println(string(jsonStr))

	// 解码的过程 jsonStr -> struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json.Unmarshal failed, err:", err)
		return
	}

	fmt.Printf("myMovie: %v\n", myMovie)
}
