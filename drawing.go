package main

import (
	"fmt"
	"log"
	"os"
	"raderh2o/tuido/todo"
	"raderh2o/tuido/tutil"

	"golang.org/x/term"
)

func (state State) drawTodos(t tutil.Terminal, todos todo.Todos) {
	t.ClearScreen()
	fmt.Println("(a)dd/(d)elete/(t)oggle/(q)uit")

	for i, todo := range todos {
		cursor := "  "

		if i == state.currentTodo {
			cursor = "->"
		}

		fmt.Printf("%v %v\n", cursor, todo)
	}
}

func (state State) addTodo(t tutil.Terminal) {
	t.ClearScreen()

	fmt.Println("Enter the body of the task you want to add :")
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		t.CloseTerminal()
		log.Fatalf("Error occured while getting the terminal size: %v", err)
	}
	t.DrawRect(tutil.Rect{X: 1, Y: 2, W: width - 1, H: 3})
	t.MoveCursorTo(3, 2)
}
