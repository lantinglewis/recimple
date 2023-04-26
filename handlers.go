package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./home.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func recipeView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a recipe with id %d", id)
}

func recipeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a recipe"))
}
