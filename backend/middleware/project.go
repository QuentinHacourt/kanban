package middleware

import (
	"encoding/json"
	"fmt"
	"kanban/database"
	"kanban/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var projectInput models.ProjectInput

	if err := json.NewDecoder(r.Body).Decode(&projectInput); err != nil {
		log.Fatalf("Unable to decode the project in the request body: %v", err)
	}

	insertID := database.InsertProject(projectInput)

	response := models.Response{
		ID:      insertID,
		Message: "Project created successfully",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int: %v", err)
	}

	project, err := database.GetProject(int64(id))

	if err != nil {
		log.Fatalf("Unable to get Project: %v", err)
	}

	if err = json.NewEncoder(w).Encode(project); err != nil {
		log.Printf("Failed to encode project: %v", err)
	}
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := database.GetAllProjects()

	if err != nil {
		log.Fatalf("Unable to get all projects: %v", err)
	}

	if err = json.NewEncoder(w).Encode(projects); err != nil {
		log.Printf("Failed to encode projects: %v", err)
	}

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the strig into int.  %v", err)
	}

	var project models.Project

	if err = json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Fatalf("Unable to decode project in the request body: %v", err)
	}

	*project.ID = id

	updatedRows := database.UpdateProject(project)

	msg := fmt.Sprintf("Project updated successfully. Total rows/record affected %v", updatedRows)

	response := models.Response{
		ID:      int64(*project.ID),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := database.DeleteProject(int64(id))

	msg := fmt.Sprintf("Project updated successfully. Total rows/record affected: %v", deletedRows)

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
