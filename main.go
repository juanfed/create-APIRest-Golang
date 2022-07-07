package main

import (
	"github.com/gorilla/mux"
)

// como luciran mis tareas y lo que responderá el cliente
// para una tarea
// Types
type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

// para lista de tareas
type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func indexRoute() {

}

func main() {
	// creo el enrutador usando esa biblioteca
	router := mux.NewRouter().StrictSlah(true) // para definir una primera ruta

	router.HandleFunc("/", indexRoute())
}
