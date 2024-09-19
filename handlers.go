package main

import (
	"html/template"
	"net/http"
)

// HomeHandler displays the to-do list and form
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := ListToDos()
	if err != nil {
		http.Error(w, "Unable to fetch to-dos", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, todos)
}

// AddToDoHandler adds a new to-do item
func AddToDoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	err := AddToDo(title)
	if err != nil {
		http.Error(w, "Unable to add to-do", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteToDoHandler deletes a to-do item
func DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := DeleteToDo(id)
	if err != nil {
		http.Error(w, "Unable to delete to-do", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// MarkCompleteHandler marks a to-do item as complete
func MarkCompleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := MarkComplete(id)
	if err != nil {
		http.Error(w, "Unable to mark to-do as complete", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// MarkAllCompleteHandler marks all to-do items as complete
func MarkAllCompleteHandler(w http.ResponseWriter, r *http.Request) {
	err := MarkAllComplete()
	if err != nil {
		http.Error(w, "Unable to mark all to-dos as complete", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
