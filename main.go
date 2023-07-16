package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
	"strings"
)

func main() {
	var title string
	var err error
	var todos []todo
	var newTodo todo

	todos, _ = readTodo()

	if os.Args[1] == "create" {
		fmt.Print("Title: ")

		reader := bufio.NewReader(os.Stdin)
		title, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("err reading title:", err)
		}

		newTodo.Title = strings.TrimRight(title, "\r\n")

		todos = append(todos, newTodo)

		jsonTodo, err := json.MarshalIndent(todos, "", " ")
		if err != nil {
			fmt.Println("err marshaling:", err)
		}

		os.WriteFile("todos.json", []byte(jsonTodo), 0644)
		fmt.Println(string(jsonTodo))
	}
}

func readTodo() ([]todo, error) {
	jsonTodo, err := os.ReadFile("todos.json")
	if err != nil {
		fmt.Println("err reading todos.json:", err)
		return []todo{}, err
	}
	var oldTodo []todo
	err = json.Unmarshal(jsonTodo, &oldTodo)
	if err != nil && len(jsonTodo) != 0{
		fmt.Println("err at unmarshal:", err)
		return []todo{}, err
	}
	return oldTodo, nil
}

type todo struct {
	Title string `json:"title"`
	Time string `json:"time"`
	Done bool `json:"done"`
}

