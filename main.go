package main

import (
	"crud/entity"
	"crud/repository"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	repo repository.TaskRepository = repository.NewTaskRepository()
)

func getTasks(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the tasks"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func createTask(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var task entity.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}
	task.ID = rand.Int63()
	repo.Save(&task)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(task)
}

/*func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprint(w, "ID no válido")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprint(w, "ID no válido")
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "La tarea con el id %v ha sido eliminada", taskID)
		}
	}

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask Task

	if err != nil {
		fmt.Fprintf(w, "ID no válido")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserta un dato válido")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)

			updatedTask.ID = task.ID
			tasks = append(tasks, updatedTask)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			fmt.Fprintf(w, "La tarea con el id %v ha sido actualizada", taskID)
		}
	}

}*/

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Que onda broder")
}
func main() {
	const port string = ":3000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	//router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	//router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	//router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
