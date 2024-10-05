package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"raderh2o/tuido/fileoperations"
	"raderh2o/tuido/todo"
	"raderh2o/tuido/tutil"
	// "strings"
)

func main() {
	homedir, _ := os.UserHomeDir()
	filepath := homedir + "/todo.txt"

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(filepath, []byte{}, 0666)
	}

	loaded, err := fileoperations.ReadFromFile(filepath)
	if err != nil {
		fmt.Println("Failed loading file: ", err)
		return
	}

	todos := todo.GetTodos(loaded)

	appState := State{
		currentTodo: 0,
		action:      Normal,
	}

	fmt.Println(todos)
	t, err := tutil.InitializeRaw()
	if err != nil {
		fmt.Println("Error switching the terminal to raw mode: ", err)
		return
	}

	for {
		b := make([]byte, 1)
		appState.drawTodos(t, todos)
		t.MoveCursorToHome()
		_, err = os.Stdin.Read(b)
		if err != nil {
			t.ClearScreen()
			fmt.Println("Error getting input: ", err)
			return
		}

		if b[0] == 'q' {
			t.ClearScreen()
			break
		}
		if b[0] == 'j' {
			if appState.currentTodo < len(todos)-1 {
				appState.currentTodo++
			}
		}
		if b[0] == 'k' {
			if appState.currentTodo > 0 {
				appState.currentTodo--
			}
		}
		if b[0] == 'd' {
			todos = append(todos[:appState.currentTodo], todos[appState.currentTodo+1:]...)
			fileoperations.WriteToFile(filepath, todos)
			if appState.currentTodo >= len(todos) && appState.currentTodo > 0 {
				appState.currentTodo--
			}
		}
		if b[0] == 't' {
			todos[appState.currentTodo].Done = !todos[appState.currentTodo].Done
			fileoperations.WriteToFile(filepath, todos)
		}
		if b[0] == 'a' {
			output := ""
			todoInput := make([]byte, 3)
			appState.addTodo(t)
			for {
				_, err := os.Stdin.Read(todoInput)
				if err != nil {
					t.ClearScreen()
					log.Fatalf("Error: %v", err)
				}
				if todoInput[0] == '\r' {
					break
				}
				if todoInput[0] == '\x7f' {
					t.CursorLeftN(1)
					fmt.Print(" ")
					t.CursorLeftN(1)
					output = output[:len(output)-1]
				}
				if string(todoInput) == "\033[A" || string(todoInput) == "\033[B" || string(todoInput) == "\033[C" || string(todoInput) == "\033[D" {
					continue
				}
				fmt.Printf("%c", todoInput[0])
				output += string(todoInput[0])
			}
			todos = append(todos, todo.Todo{Content: output, Done: false})
			fileoperations.WriteToFile(filepath, todos)
		}
	}

	t.CloseTerminal()

}
