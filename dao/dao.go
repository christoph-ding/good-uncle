package dao

import (
  "log"

  mgo "gopkg.in/mgo.v2"
  . "github.com/ironpup/good-uncle/models"
  // "gopkg.in/mgo.v2/bson"
)

// Data base connection object
type MealPlanDAO struct {
  Server   string
  Database string
}

var db *mgo.Database

const (
  PLANS_COLLECTION = "plans"
)

func (m *MealPlanDAO) Connect() {
  session, err := mgo.Dial(m.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(m.Database)
}

// a query for each endpoint function
// func plansAll(w http.ResponseWriter, r *http.Request) {

// func plansByMarket(w http.ResponseWriter, r *http.Request) {

// func plansByID(w http.ResponseWriter, r *http.Request) {

// Insert a movie into database
// func (m *MoviesDAO) Insert(movie Movie) error {
//   err := db.C(COLLECTION).Insert(&movie)
//   return err
// }

func (m *MealPlanDAO) plansCreate(plan Plan) error  {
  err := db.C(PLANS_COLLECTION).Insert(&plan)
  return err
}

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
