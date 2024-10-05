package main

type Action int

const (
	AddTodo Action = iota
	RemoveTodo
	ToggleTodo
	ChangeTodo
	Normal
	Quit
)

type State struct {
	currentTodo int
	action      Action
}
