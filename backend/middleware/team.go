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

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var teamInput models.TeamInput

	if err := json.NewDecoder(r.Body).Decode(&teamInput); err != nil {
		log.Fatalf("Unable to decode the team in the request body: %v", err)
	}

	insertID := database.InsertTeam(teamInput)

	response := models.Response{
		ID:      insertID,
		Message: "Team created successfully",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int: %v", err)
	}

	team, err := database.GetTeam(int64(id))

	if err != nil {
		log.Fatalf("Unable to get Team: %v", err)
	}

	if err = json.NewEncoder(w).Encode(team); err != nil {
		log.Printf("Failed to encode team: %v", err)
	}
}

func GetAllTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	teams, err := database.GetAllTeams()

	if err != nil {
		log.Fatalf("Unable to get all teams: %v", err)
	}

	if err = json.NewEncoder(w).Encode(teams); err != nil {
		log.Printf("Failed to encode teams: %v", err)
	}

}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the strig into int.  %v", err)
	}

	var team models.Team

	if err = json.NewDecoder(r.Body).Decode(&team); err != nil {
		log.Fatalf("Unable to decode team in the request body: %v", err)
	}

	*team.ID = id

	updatedRows := database.UpdateTeam(team)

	msg := fmt.Sprintf("Team updated successfully. Total rows/record affected %v", updatedRows)

	response := models.Response{
		ID:      int64(*team.ID),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := database.DeleteTeam(int64(id))

	msg := fmt.Sprintf("Team updated successfully. Total rows/record affected: %v", deletedRows)

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
