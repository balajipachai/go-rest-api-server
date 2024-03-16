package library

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

// The `func (a *App) Initialize()` method is initializing the application by performing the following
// tasks:
// 1. Opening a connection to a SQLite database located at "../practiceit.db".
// 2. If there is an error opening the database connection, it logs the error and exits the
// application.
// 3. Creates a new Gorilla Mux router for handling HTTP requests.
// 4. Calls the `initializeRoutes()` method to set up the routes for the application.
func (a *App) Initialize() {
	var err error
	a.DB, err = sql.Open("sqlite3", "../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// The `initializeRoutes()` method in the provided Go code snippet is setting up a route in the Gorilla
// Mux router for handling HTTP GET requests to the root path ("/").
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", helloWorld).Methods("GET")
}

// The helloWorld function in Go writes a specific message to the http.ResponseWriter.
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Balaji Pachai, is the wisest blockchain developer\n")
}

// The `func (a *App) Run()` method in the provided Go code snippet is responsible for starting the
// HTTP server and listening for incoming requests on the specified port. Here's what it does:
func (a *App) Run() {
	fmt.Println("Server started and listening on: ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
