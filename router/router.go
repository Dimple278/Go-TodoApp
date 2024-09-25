package router

import (
	"net/http"

	"github.com/Dimple278/Go-TodoApp/controller"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controller.HomeHandler)
	http.HandleFunc("/todos", controller.AddToDoHandler)
	http.HandleFunc("/todos/", controller.DeleteToDoHandler)
	http.HandleFunc("/todos/complete/", controller.MarkCompleteHandler)
	http.HandleFunc("/todos/complete-all", controller.MarkAllCompleteHandler)
}
