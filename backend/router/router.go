package router

import (
	"kanban/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Allow CORS
	router.Use(middleware.CORS)

	// ===========
	// =   API   =
	// ===========

	// stories
	router.HandleFunc("/story/{id}", middleware.GetStory).Methods("GET", "OPTIONS")
	router.HandleFunc("/story", middleware.GetAllStories).Methods("GET", "OPTIONS")
	router.HandleFunc("/story", middleware.CreateStory).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/story/{id}", middleware.UpdateStory).Methods("PUT", "OPTIONS")
	router.HandleFunc("/story/{id}", middleware.DeleteStory).Methods("DELETE", "OPTIONS")

	// developers
	router.HandleFunc("/developer/{id}", middleware.GetDeveloper).Methods("GET", "OPTIONS")
	router.HandleFunc("/developer", middleware.GetAllDevelopers).Methods("GET", "OPTIONS")
	router.HandleFunc("/developer", middleware.CreateDeveloper).Methods("POST", "OPTIONS")
	router.HandleFunc("/developer/{id}", middleware.UpdateDeveloper).Methods("PUT", "OPTIONS")
	router.HandleFunc("/developer/{id}", middleware.DeleteDeveloper).Methods("DELETE", "OPTIONS")

	// projecs
	router.HandleFunc("/project/{id}", middleware.GetProject).Methods("GET", "OPTIONS")
	router.HandleFunc("/project", middleware.GetAllProjects).Methods("GET", "OPTIONS")
	router.HandleFunc("/project", middleware.CreateProject).Methods("POST", "OPTIONS")
	router.HandleFunc("/project/{id}", middleware.UpdateProject).Methods("PUT", "OPTIONS")
	router.HandleFunc("/project/{id}", middleware.DeleteProject).Methods("DELETE", "OPTIONS")

	// teams
	router.HandleFunc("/team/{id}", middleware.GetTeam).Methods("GET", "OPTIONS")
	router.HandleFunc("/team", middleware.GetAllTeams).Methods("GET", "OPTIONS")
	router.HandleFunc("/team", middleware.CreateTeam).Methods("POST", "OPTIONS")
	router.HandleFunc("/team/{id}", middleware.UpdateTeam).Methods("PUT", "OPTIONS")
	router.HandleFunc("/team/{id}", middleware.DeleteTeam).Methods("DELETE", "OPTIONS")

	return router
}
