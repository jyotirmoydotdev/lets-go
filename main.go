package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
  if r.Method != http.MethodGet {
    w.Header().Set("Allowed", http.MethodGet)
    w.WriteHeader(http.StatusMethodNotAllowed)
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    log.Println("405 - /snippet/view - Method Not Allowed")
    return
  }

  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }
  log.Println("200 - /snippet/view")
  fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
  if r.Method != "POST" {
    w.Header().Set("Allow", http.MethodPost)
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method Not Allowed"))
    http.Error(w,"Method Not Allowed", http.StatusMethodNotAllowed)
    log.Println("405 - /snippet/create - Method Not Allowed")
    return 
  }
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
