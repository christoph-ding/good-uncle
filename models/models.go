package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
)

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