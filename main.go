package main

import (
	"KaiJi-Casino/internal/pkg/configs"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"encoding/json"
	"fmt"
)

func main() {
	testing()
	configs.New()
	db.New()
}

func testing() {
	a := collection.SportsData{}
	//var sd collection.SportsData
	//json.Unmarshal([]byte(a), &sd)
	b, _ := json.Marshal(a)
	c := string(b)
	fmt.Println(string(c))
}
