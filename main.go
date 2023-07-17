package main

import (
	"fmt"
	"os"
	"encoding/json"
	"strings"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	var todos []todo
	var newTodo todo

	todos, _ = readTodo()

	switch os.Args[1] {
	case "create":
		var title string
		title = os.Args[2]

		newTodo.Title = strings.TrimRight(title, "\r\n")

		todos = append(todos, newTodo)

		updateJsonTodo(todos)

		fmt.Printf("%s added to the list\n", title)
		showTable(todos)

	case "settime":
		// todo: change datatype to time.Time
		var date string
		var id int

		id, _ = strconv.Atoi(os.Args[2])
		date = os.Args[3]
		todos[id].Time = date

		updateJsonTodo(todos)
		showTable(todos)
		
	case "delete":
		var id int
		var newTodo []todo
		id, _ = strconv.Atoi(os.Args[2])
		if id < len(todos) {
			newTodo = append(todos[:id], todos[id+1:]...)
			updateJsonTodo(newTodo)
			showTable(newTodo)
		} else {
			fmt.Println("Out of bounds")
		}
	case "done":
		var id int
		id, _ = strconv.Atoi(os.Args[2])

		todos[id].Done = true
		updateJsonTodo(todos)
		showTable(todos)
	case "show":
		showTable(todos)

	default:
		fmt.Println("Invalid command")
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

func showTable(todos []todo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Time", "Done"})
	for i, e := range todos {
		t.AppendRow([]interface{}{ i, e.Title, e.Time, e.Done })
	}
	t.SetStyle(table.StyleColoredGreenWhiteOnBlack)
	t.Render()
}

func updateJsonTodo(todos []todo) {
		jsonTodo, err := json.MarshalIndent(todos, "", " ")
		if err != nil {
			fmt.Println("err marshaling:", err)
		}
		os.WriteFile("todos.json", jsonTodo, 0644)
}
