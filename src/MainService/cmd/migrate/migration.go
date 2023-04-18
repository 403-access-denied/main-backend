package main

import (
	"github.com/403-access-denied/main-backend/src/MainService/DBConfiguration"
	Model2 "github.com/403-access-denied/main-backend/src/MainService/Model"
)

func main() {
	Migration()
}

func Migration() {
	var models []interface{}
	// add the model to migrate the database
	models = append(models, &Model2.User{})
	models = append(models, &Model2.Poster{})
	models = append(models, &Model2.Category{})
	models = append(models, &Model2.Conversation{})
	models = append(models, &Model2.Message{})
	models = append(models, &Model2.Image{})
	models = append(models, &Model2.Address{})
	models = append(models, &Model2.MarkedPoster{})
	DBConfiguration.InitDB()
	DBConfiguration.MigrateModel(models)
	DBConfiguration.CloseDB()
}
