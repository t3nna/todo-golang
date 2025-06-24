package main

import (
	"fmt"
	"log"
)

func main() {
	opts, err := getOpts()
	fmt.Println("")
	if err != nil {
		log.Fatal("error: %v", err)
	}
	fmt.Println(opts)
	todo := NewTodo("./data/data.json")
	data := todo.GetValueAll()
	fmt.Println(data)

}
