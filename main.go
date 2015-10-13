package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/cats", CatIndex)
	router.HandleFunc("/cats/{catFactID}", CatFactShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index ayyyy
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Cat Lover, %q", html.EscapeString(r.URL.Path))
}

func CatIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cat Index!")
}

func CatFactShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["catFactID"]
	fmt.Fprintln(w, "Cat Fact:", todoId)
}
