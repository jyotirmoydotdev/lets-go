package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    log.Println("404 -", r.URL.Path)
    return
  }
  log.Println("200 - /")
  w.Write([]byte("Hello from JB"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
  log.Println("200 - /snippet/view")
  w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
  log.Println("200 - /snippet/create")
  w.Write([]byte("Create a new snippet..."))
}

func main(){
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

  log.Println("Sarting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
