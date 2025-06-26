package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
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

func (todo *Todo) SetValue(item *TodoItem, key ...string) {
	if len(key) > 0 {
		todo.data.Todo[key[0]] = *item
		return
	}
	newId := uuid.New().String()
	item.ID = newId
	todo.data.Todo[newId] = *item

}

func (todo *Todo) DeleteValue(key string) {
	_, found := todo.GetValueById(key)
	if found == false {
		log.Fatal("Something went wrong")
		return
	}

	if found {
		delete(todo.data.Todo, key)

	}
	todo.Save()
}

func (todo *Todo) MakeComplete(key string) {
	item, found := todo.GetValueById(key)

	if found == false {
		log.Fatal("Something went wrong", item)
		return
	}
	item.IsComplete = true
	todo.SetValue(item, item.ID)
	todo.Save()

}

func defaultTodo(path string) *Todo {
	return &Todo{
		data: &Data{map[string]TodoItem{}},
		path: path,
	}
}

func (todo *Todo) Save() error {
	dir := path.Dir(todo.path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}

	}
	jsonString, err := json.Marshal(todo.data)
	if err != nil {
		return err
	}
	err = os.WriteFile(todo.path, jsonString, 0755)
	if err != nil {
		return err
	}

	return nil
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
