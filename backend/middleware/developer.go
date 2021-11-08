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
	var developerInput models.DeveloperInput

	if err := json.NewDecoder(r.Body).Decode(&developerInput); err != nil {
		log.Fatalf("Unable to decode the developer in the request body: %v", err)
	}

	insertID, err := database.InsertDeveloper(developerInput)
	if err != nil {
		errorHandler(w, err)
		return
	}

	response := models.Response{
		ID:      insertID,
		Message: "Developer created successfully",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func GetDeveloper(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		errorHandler(w, fmt.Errorf("Unable to convert the string into int: %v", err))
	}

	developer, err := database.GetDeveloper(int64(id))

	if err != nil {
		errorHandler(w, err)
	}

	if err = json.NewEncoder(w).Encode(developer); err != nil {
		errorHandler(w, fmt.Errorf("Failed to encode developer: %v", err))
	}
}

func GetAllDevelopers(w http.ResponseWriter, r *http.Request) {
	developers, err := database.GetAllDevelopers()

	if err != nil {
		errorHandler(w, fmt.Errorf("Unable to get all developers: %v", err))
	}

	if err = json.NewEncoder(w).Encode(developers); err != nil {
		errorHandler(w, fmt.Errorf("Failed to encode developers: %v", err))
	}
}

func UpdateDeveloper(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		errorHandler(w, fmt.Errorf("Unable to convert the strig into int.  %v", err))
	}

	var developer models.Developer

	if err = json.NewDecoder(r.Body).Decode(&developer); err != nil {
		errorHandler(w, fmt.Errorf("Unable to decode developer in the request body: %v", err))
	}

	*developer.ID = id

	updatedRows, err := database.UpdateDeveloper(developer)
	if err != nil {
		errorHandler(w, err)
	}

	msg := fmt.Sprintf("Developer updated successfully. Total rows/record affected %v", updatedRows)

	response := models.Response{
		ID:      int64(*developer.ID),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		errorHandler(w, fmt.Errorf("Failed to encode response: %v", err))
	}
}

func DeleteDeveloper(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows, err := database.DeleteDeveloper(int64(id))
	if err != nil {
		errorHandler(w, err)
	}

	msg := fmt.Sprintf("Developer updated successfully. Total rows/record affected: %v", deletedRows)

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		errorHandler(w, fmt.Errorf("Failed to encode response: %v", err))
	}
}

func errorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal: %s", err)
	}
	w.Write(jsonResp)
	return

}
