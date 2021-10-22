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

func CreateDeveloper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var developerInput models.DeveloperInput

	if err := json.NewDecoder(r.Body).Decode(&developerInput); err != nil {
		log.Fatalf("Unable to decode the developer in the request body: %v", err)
	}

	insertID := database.InsertDeveloper(developerInput)

	response := models.Response{
		ID:      insertID,
		Message: "Developer created successfully",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func GetDeveloper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int: %v", err)
	}

	developer, err := database.GetDeveloper(int64(id))

	if err != nil {
		log.Fatalf("Unable to get Developer: %v", err)
	}

	if err = json.NewEncoder(w).Encode(developer); err != nil {
		log.Printf("Failed to encode developer: %v", err)
	}
}

func GetAllDevelopers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	developers, err := database.GetAllDevelopers()

	if err != nil {
		log.Fatalf("Unable to get all developers: %v", err)
	}

	if err = json.NewEncoder(w).Encode(developers); err != nil {
		log.Printf("Failed to encode developers: %v", err)
	}

}

func UpdateDeveloper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the strig into int.  %v", err)
	}

	var developer models.Developer

	if err = json.NewDecoder(r.Body).Decode(&developer); err != nil {
		log.Fatalf("Unable to decode developer in the request body: %v", err)
	}

	*developer.ID = id

	updatedRows := database.UpdateDeveloper(developer)

	msg := fmt.Sprintf("Developer updated successfully. Total rows/record affected %v", updatedRows)

	response := models.Response{
		ID:      int64(*developer.ID),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func DeleteDeveloper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := database.DeleteDeveloper(int64(id))

	msg := fmt.Sprintf("Developer updated successfully. Total rows/record affected: %v", deletedRows)

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
