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
	router.HandleFunc("/story/{id}", middleware.GetStory).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/story", middleware.GetAllStories).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/story", middleware.CreateStory).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/story/{id}", middleware.UpdateStory).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/story/{id}", middleware.DeleteStory).Methods(http.MethodDelete, http.MethodOptions)

	// developers
	router.HandleFunc("/developer/{id}", middleware.GetDeveloper).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/developer", middleware.GetAllDevelopers).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/developer", middleware.CreateDeveloper).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/developer/{id}", middleware.UpdateDeveloper).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/developer/{id}", middleware.DeleteDeveloper).Methods(http.MethodDelete, http.MethodOptions)

	// projecs
	router.HandleFunc("/project/{id}", middleware.GetProject).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/project", middleware.GetAllProjects).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/project", middleware.CreateProject).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/project/{id}", middleware.UpdateProject).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/project/{id}", middleware.DeleteProject).Methods(http.MethodDelete, http.MethodOptions)

	// teams
	router.HandleFunc("/team/{id}", middleware.GetTeam).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/team", middleware.GetAllTeams).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/team", middleware.CreateTeam).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/team/{id}", middleware.UpdateTeam).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/team/{id}", middleware.DeleteTeam).Methods(http.MethodDelete, http.MethodOptions)

	return router
}
