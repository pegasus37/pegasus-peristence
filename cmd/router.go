package main

import (
	"github.com/gorilla/mux"
)

const (
	studentList = "/api/v1/student/"
)

func setUpRouter(router *mux.Router) {
	router.HandleFunc(studentList, handler.GetStudentList).Methods("GET")
	router.HandleFunc(studentList, handler.PostStudentList).Methods("POST")
}
