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
	todo.SetValue(&TodoItem{
		ID:          "1233",
		Description: "Pick up package",
		CreatedAt:   "some date",
		IsComplete:  false,
	})
	err = todo.Save()
	if err != nil {
		log.Fatalf("Unable to write file: %v", err)
	}
	fmt.Println(data)

}
