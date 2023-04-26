package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", home)
	router.HandlerFunc(http.MethodGet, "/recipe/view/:id", recipeView)
	router.HandlerFunc(http.MethodPost, "/recipe/create", recipeCreate)

	http.ListenAndServe(":4000", router)
}
