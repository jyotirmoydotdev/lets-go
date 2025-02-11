package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Println("404 - / - Path Not Found")
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println(http.StatusMethodNotAllowed, "- /snippet/view - Method Not Allow")
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		log.Println(http.StatusNotFound, "- /snippet/view?id= - Id is missing")
		return
	}

	log.Printf("%d - /snippet/view?id=%d - All Ok", http.StatusOK, id)
	fmt.Fprintf(w, "display a specific snippet of ID: %d....", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("%d - /snippet/create - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("%d - /snippet/create - All OK", http.StatusOK)
	w.Write([]byte("Create a new snippet...."))
}
