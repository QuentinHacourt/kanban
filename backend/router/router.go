package router

import (
	"kanban/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// ===========
	// =   API   =
	// ===========

	// stories
	router.HandleFunc("/story/{id}", middleware.GetStory).Methods("GET", "OPTIONS")
	router.HandleFunc("/story", middleware.GetAllStories).Methods("GET", "OPTIONS")
	router.HandleFunc("/story", middleware.CreateStory).Methods("POST", "OPTIONS")
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

	// Allow CORS
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
