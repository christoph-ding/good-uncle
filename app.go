package main

import (
  "fmt"
  "net/http"
  "log"

  "github.com/gorilla/mux"
)

// Endpoints by feature
// I omit adding the descriptor 'handler' or 'endpoint' here because I thought 
// that it adds unneccesary visual noise - it should be clear by how these functions
// are used that they are handlers for requests

// Plans
func plansAll(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "getting all plans")
}

func plansByMarket(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "getting all plans by market")
}

func plansByID(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "getting all plans by id")
}

func plansCreate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "creating plan")
}

func plansUpdate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "updating plan")
}

func plansDelete(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "deleting plan")
}

// create server
func main() {
  fmt.Println("creating server ...")

  r := mux.NewRouter()

  r.HandleFunc("/plans", plansAll).Methods("GET")
  if err := http.ListenAndServe(":3000", r); err != nil {
    log.Fatal(err)
  }
}
