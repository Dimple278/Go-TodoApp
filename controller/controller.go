package controller

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/Dimple278/Go-TodoApp/models"
	"github.com/Dimple278/Go-TodoApp/utils"
)

// HomeHandler displays the to-do list and form
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := models.ListToDos()
	if err != nil {
		utils.ErrorHandler(w, "Unable to fetch to-dos", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		utils.ErrorHandler(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, todos)
}

func AddToDoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorHandler(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		utils.ErrorHandler(w, "Title is required", http.StatusBadRequest)
		return
	}

	err := models.AddToDo(title)
	if err != nil {
		utils.ErrorHandler(w, "Unable to add to-do", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	method := r.URL.Query().Get("_method")
	if method != "DELETE" {
		utils.ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	if !strings.HasPrefix(path, "/todos/") {
		utils.ErrorHandler(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(path, "/todos/")
	if id == "" {
		utils.ErrorHandler(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := models.DeleteToDo(id)
	if err != nil {
		utils.ErrorHandler(w, "Unable to delete to-do", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MarkCompleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ErrorHandler(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := models.MarkComplete(id)
	if err != nil {
		utils.ErrorHandler(w, "Unable to mark to-do as complete", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MarkAllCompleteHandler(w http.ResponseWriter, r *http.Request) {
	err := models.MarkAllComplete()
	if err != nil {
		utils.ErrorHandler(w, "Unable to mark all to-dos as complete", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
