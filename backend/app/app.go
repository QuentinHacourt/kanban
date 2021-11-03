package app

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
)

type App struct {
	Database *sql.DB
	Router   *mux.Router
}

func (app *App) start() {
	// app.router = router.Router()

	fmt.Println("starting server on port 8080...")

	// log.Fatal(http.ListenAndServe(":8080", app.router))
}
