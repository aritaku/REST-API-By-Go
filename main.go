package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Todo is todo list data
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is todo list slice
type Todos []Todo

// Index is なんぞ
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!")
}

// TodoIndex is
func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Todo Index!")
}

// TodoShow is
func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoid"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)
	router.GET("/todos/:todoId", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
