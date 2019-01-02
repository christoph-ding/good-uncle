package dao

import (
  "log"

  "github.com/mlabouardy/movies-restapi/models"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

// Data base connection object
type MealPlanDAO struct {
  Server   string
  Database string
}

var db *mgo.Database

// CONSTANTS - plans, users <- should probably be moved out <- this file should ONLY give the dao

// a query for each endpoint function
// func plansAll(w http.ResponseWriter, r *http.Request) {

// func plansByMarket(w http.ResponseWriter, r *http.Request) {

// func plansByID(w http.ResponseWriter, r *http.Request) {

// func plansCreate(w http.ResponseWriter, r *http.Request) {

// func plansUpdate(w http.ResponseWriter, r *http.Request) {

// func plansDelete(w http.ResponseWriter, r *http.Request) {

// func planDuplicate(w http.ResponseWriter, r *http.Request) {

// func userSignUp(w http.ResponseWriter, r *http.Request) {

// func userAddMealPlan(w http.ResponseWriter, r *http.Request) {

// func (m *MealPlanDAO) Connect() {
//   session, err := mgo.Dial(m.Server)
//   if err != nil {
//     log.Fatal(err)
//   }
//   db = session.DB(m.Database)
// }
