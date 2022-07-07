package main

import (
	"fmt"
	"log"
	"net/http"

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

// con esto podré empezar a reponder algo
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esto esta funcionando de maravilla")
}

func main() {
	// creo el enrutador usando esa biblioteca
	router := mux.NewRouter().StrictSlash(true) // para definir una primera ruta

	router.HandleFunc("/", indexRoute)

	// creacion del servidor http que tendra el puerto donde escucha y el enrutador
	log.Fatal(http.ListenAndServe(":3000", router))
}
