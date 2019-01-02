package main

import (
  "fmt"
  "net/http"
  "log"

  "gopkg.in/mgo.v2/bson"
  "time"

  "github.com/gorilla/mux"
)

// Models

// Each plan will have a name, weekly cost, spring semester start and end dates, and fall semester start and end dates.  Each plan will be associated with one and only one of the markets in which we operate.  To help educate customers about our meal plans, internal users will also need to be able to enter three pieces of marketing text that will be displayed in our customer-facing apps, including our mobile apps.  In addition, some of the schools where we operate have trimesters rather than semesters, so internal users will need to have the option to enter the start and end date for a third academic period.

type Plan struct {
  ID            bson.ObjectId `bson:"_id" json:"id"`
  name          string        `bson:"name" json:"name"`
  cost          float64       `bson:"cost" json:"cost"`
  springStart   time.Time     `bson:"springStart" json:"springStart"`
  springEnd     time.Time     `bson:"springEnd" json:"springEnd"`
  fallStart     time.Time     `bson:"fallStart" json:"fallStart"`
  fallEnd       time.Time     `bson:"fallEnd" json:"fallEnd"`
  thirdStart    time.Time     `bson:"thirdStart" json:"thirdStart"`
  thirdEnd      time.Time     `bson:"thirdEnd" json:"thirdEnd"`
  market        string        `bson:"market" json:"market"`
  textOne       string        `bson:"textOne" json:"textOne"`
  textTwo       string        `bson:"textTwo" json:"textTwo"`
  textThree     string        `bson:"textThree" json:"textThree"`
}

type User struct {
  ID            bson.ObjectId `bson:"_id" json:"id"`
  mealPlanId    []string      `bson:"mealPlanId" json:"mealPlanId"`
}

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
