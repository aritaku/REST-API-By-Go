package main

import "time"

// Todo is todo list data
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is todo list slice
type Todos []Todo
