package main

import (
  "fmt"
  "net/http"
  "log"

  "gopkg.in/mgo.v2/bson"
  

  "encoding/json"
  "github.com/gorilla/mux"
  . "github.com/ironpup/good-uncle/dao"
  . "github.com/ironpup/good-uncle/models"
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
  fmt.Fprintln(w, "creating plan with request")
  fmt.Fprintln(w, r)

  defer r.Body.Close()
  var plan Plan

  if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }

  plan.ID = bson.NewObjectId()

  if err := dao.planCreate(movie); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJson(w, http.StatusCreated, plan)
}

func plansUpdate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "updating plan")
}

func plansDelete(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "deleting plan")
}

func planDuplicate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "duplicating plan")
}

// Users
func userSignUp(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "creating user")
}

func userAddMealPlan(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "adding meal plan to user")
}

// helper functions
func respondWithError(w http.ResponseWriter, code int, msg string) {
  respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

// create server
func main() {
  fmt.Println("creating server ...")

  r := mux.NewRouter()

  r.HandleFunc("/plans", plansAll).Methods("GET")
  r.HandleFunc("/plans/market/{id}", plansByMarket).Methods("GET")
  r.HandleFunc("/plans/id/{id}", plansByID).Methods("GET")
  r.HandleFunc("/plans", plansCreate).Methods("POST")
  r.HandleFunc("/plans", plansUpdate).Methods("PUT")
  r.HandleFunc("/plans", plansDelete).Methods("DELETE")
  r.HandleFunc("/plans/duplicate", planDuplicate).Methods("POST")
  r.HandleFunc("/user", userSignUp).Methods("POST")
  r.HandleFunc("/user/add-plan", userAddMealPlan).Methods("PUT")

  if err := http.ListenAndServe(":3000", r); err != nil {
    log.Fatal(err)
  }
}
