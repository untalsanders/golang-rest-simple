package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
	Content string `json:"Content"`
}

type tasks []task

var myTasks = tasks {
	{
		ID: 1,
		Name: "Task one",
		Content: "Some task",
	},
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API made with Go Lang")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myTasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, task := range myTasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(myTasks) + 1
	myTasks = append(myTasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, task := range myTasks {
		if task.ID == taskID {
			myTasks = append(myTasks[:i], myTasks[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove succesfully", taskID)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	var updatedTask task
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please, Enter a valid task")
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range myTasks {
		if task.ID == taskID {
			myTasks = append(myTasks[:i], myTasks[i+1:]...)
			updatedTask.ID = taskID
			myTasks = append(myTasks, updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated succesfully", taskID)
		}
	}
}

func RunServer() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexRoute)
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3500", router))
}