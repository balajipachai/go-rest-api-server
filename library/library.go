package library

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB   *sql.DB
	Port string
}

func (a *App) Initialize() {
	var err error
	a.DB, err = sql.Open("sqlite3", "../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Balaji Pachai, is the wisest blockchain developer\n")
}

func (a *App) Run() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on: ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))

}
