package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TodoItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	IsComplete  bool   `json:"isComplete"`
}

type Data struct {
	Todo map[string]TodoItem `json:"todo"`
}
type Todo struct {
	data *Data
	path string
}

func (todo *Todo) GetValueById(key string) (item *TodoItem, found bool) {
	if item, ok := todo.data.Todo[key]; ok == true {
		return &item, true
	}
	return nil, false
}

func (todo *Todo) GetValueAll() Data {
	return *todo.data
}

func defaultTodo(path string) *Todo {
	return &Todo{
		data: &Data{map[string]TodoItem{}},
		path: path,
	}
}

// if we have file in place
// read file and return To do
// if no file create new To do

func NewTodo(path string) *Todo {
	if _, err := os.Stat(path); err == nil {
		contents, err := os.ReadFile(path)
		if err != nil {
			return defaultTodo(path)
		}
		var data Data
		err = json.Unmarshal(contents, &data)
		if err != nil {
			return defaultTodo(path)
		}
		fmt.Println(data)
		return &Todo{
			data: &data,
			path: path,
		}

	}
	return defaultTodo(path)
}
