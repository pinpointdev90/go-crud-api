package main

import (
	"net/http"

	"github.com/pinpointdev90/go-crud-api/tree/main/controller"
	"github.com/pinpointdev90/go-crud-api/tree/main/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
