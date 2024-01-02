package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var user_action []string

main_loop:
	for {
		fmt.Print("Type 'add', 'show', 'complete', 'edit' 'exit': ")

		// Read user input
		user_action = sliceInput()

		if len(user_action) == 0 {
			fmt.Println("Invalid input provided")
			continue main_loop
		}

		switch user_action[0] {
		case "add":
			{
				addTodos(user_action[1:])
				showTodos()
			}
		case "show":
			showTodos()
		case "complete":
			{
				completeTodos(user_action)
				showTodos()
			}
		case "edit":
			{
				editTodos(user_action)
				showTodos()
			}
		case "exit":
			{
				fmt.Println("Good bye!")
				break main_loop
			}
		default:
			fmt.Println("Invalid input provided")
		}
	}
}

func addTodos(value []string) {
	todo := sliceToString(value)
	todo = strings.Trim(todo, " ")

	todos := readTodosFile()
	todos = append(todos, todo)

	writeTodosFile(todos)
}

func showTodos() {
	todos := readTodosFile()

	if len(todos) == 0 {
		fmt.Println("No items in the todo list.")
		return
	}

	for i := 0; i < len(todos); i++ {
		fmt.Printf("%d - %s\n", i+1, (todos)[i])
	}

}

func completeTodos(userInput []string) {
	num, err := strconv.Atoi(userInput[1])

	if err != nil {
		fmt.Println("Invalid index provided.")
		return
	}

	todos := readTodosFile()

	todos = append(todos[:num-1], todos[num:]...)
	writeTodosFile(todos)
}

func editTodos(userInput []string) {
	todo := sliceToString(userInput[2:])
	todo = strings.Trim(todo, " ")

	num, err := strconv.Atoi(userInput[1])

	if err != nil {
		fmt.Println("Invalid index provided.")
		return
	}

	todos := readTodosFile()
	todos[num-1] = todo
	writeTodosFile(todos)
}

/*
Get user input as a string and convert it into slice and return the slice
*/
func sliceInput() []string {
	var inputLine []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine = append(inputLine, strings.Fields(scanner.Text())...)
	return inputLine
}

func sliceToString(value []string) string {
	todo := ""
	for _, value := range value {
		todo = todo + " " + value
	}

	return todo
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readTodosFile() []string {
	data, err := os.ReadFile("./todos.txt")
	check(err)

	if len(data) == 0 {
		return nil
	}

	todos := []string{}
	strTodo := ""

outer_loop:
	for _, value := range data {

		if value == 10 {
			todos = append(todos, strTodo)
			strTodo = ""
			continue outer_loop
		}

		strTodo = strTodo + string(value)
	}

	if len(strTodo) > 0 {
		todos = append(todos, strTodo)
	}

	return todos

}

func writeTodosFile(todos []string) {
	f, err := os.Create("./todos.txt")
	check(err)

	defer f.Close()

	for _, value := range todos {
		f.WriteString(value + "\n")
	}

	f.Sync()
}
