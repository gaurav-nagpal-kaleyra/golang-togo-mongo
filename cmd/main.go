package main

import (
	"context"
	"fmt"
	"golang-todo/config"
	"golang-todo/constants"
	"golang-todo/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	defer func() {
		_ = config.MongoConn.Disconnect(context.Background())
	}()

	// initialize mongodb connection
	if err := config.ConnectMongo(); err != nil {
		fmt.Println("error connecting to mongodb")
		return
	}

	router := mux.NewRouter()

	// handle CRUD operations
	router.HandleFunc(constants.Task, usecase.CreateTaskHandler)
	router.HandleFunc(constants.Task, usecase.DeleteTaskHandler)
	router.HandleFunc(constants.Task, usecase.UpdateTaskHandler)
	router.HandleFunc(constants.Task, usecase.GetTaskHandler)

	// handle if invalid endpoint
	router.HandleFunc("/", usecase.AllPathHandler)

	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println("error starting server", err.Error())
		return
	}

}
