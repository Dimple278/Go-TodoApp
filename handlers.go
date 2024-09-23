package main

import (
	"html/template"
	"net/http"
	"strings"
)

// HomeHandler displays the to-do list and form
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := ListToDos()
	if err != nil {
		ErrorHandler(w, "Unable to fetch to-dos", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		ErrorHandler(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, todos)
}

func AddToDoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		ErrorHandler(w, "Title is required", http.StatusBadRequest)
		return
	}

	err := AddToDo(title)
	if err != nil {
		ErrorHandler(w, "Unable to add to-do", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	method := r.URL.Query().Get("_method")
	if method != "DELETE" {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	if !strings.HasPrefix(path, "/todos/") {
		ErrorHandler(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(path, "/todos/")
	if id == "" {
		ErrorHandler(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := DeleteToDo(id)
	if err != nil {
		ErrorHandler(w, "Unable to delete to-do", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MarkCompleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		ErrorHandler(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := MarkComplete(id)
	if err != nil {
		ErrorHandler(w, "Unable to mark to-do as complete", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MarkAllCompleteHandler(w http.ResponseWriter, r *http.Request) {
	err := MarkAllComplete()
	if err != nil {
		ErrorHandler(w, "Unable to mark all to-dos as complete", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
