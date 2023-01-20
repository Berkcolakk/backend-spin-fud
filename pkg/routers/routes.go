package routes

import (
	"net/http"
	person "spin-fud/pkg/controllers"

	"github.com/gorilla/mux"
)

func GenerateRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/people", person.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", person.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", person.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", person.DeletePersonEndpoint).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
