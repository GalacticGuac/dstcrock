package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.cookFood).Methods("GET", "POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	// Notice no DB code here

	fmt.Println("Server here")

	a := App{}
	a.Initialize()

	a.Run(":8010")

	// server, err := portal.NewServer()
	// if err != nil {
	// 	panic(err)
	// }

	// err = server.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func (a *App) cookFood(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Result here")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
