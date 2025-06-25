package main

import (
	"log"
)

func main() {
	opts, err := getOpts()
	if err != nil {
		log.Fatal("error: %v", err)
	}
	todo := NewTodo("./data/data.json")

	err = handleOptions(opts, todo)

	if err != nil {
		log.Fatalf("Unable to write file: %v", err)
	}

}
