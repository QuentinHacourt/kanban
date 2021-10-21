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

func CreateStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var storyInput models.StoryInput

	if err := json.NewDecoder(r.Body).Decode(&storyInput); err != nil {
		log.Fatalf("Unable to decode the story in the request body: %v", err)
	}

	insertID := database.InsertStory(storyInput)

	response := models.Response{
		ID:      insertID,
		Message: "Story created successfully",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func GetStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int: %v", err)
	}

	story, err := database.GetStory(int64(id))

	if err != nil {
		log.Fatalf("Unable to get Story: %v", err)
	}

	if err = json.NewEncoder(w).Encode(story); err != nil {
		log.Printf("Failed to encode story: %v", err)
	}
}

func GetAllStories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	stories, err := database.GetAllStories()

	if err != nil {
		log.Fatalf("Unable to get all stories: %v", err)
	}

	if err = json.NewEncoder(w).Encode(stories); err != nil {
		log.Printf("Failed to encode stories: %v", err)
	}

}

func UpdateStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the strig into int.  %v", err)
	}

	var story models.Story

	if err = json.NewDecoder(r.Body).Decode(&story); err != nil {
		log.Fatalf("Unable to decode story in the request body: %v", err)
	}

	story.ID = &id

	updatedRows := database.UpdateStory(story)

	msg := fmt.Sprintf("Story updated successfully. Total rows/record affected %v", updatedRows)

	response := models.Response{
		ID:      int64(*story.ID),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func DeleteStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := database.DeleteStory(int64(id))

	msg := fmt.Sprintf("Story updated successfully. Total rows/record affected: %v", deletedRows)

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
