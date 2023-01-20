package routes

import (
	"net/http"
	authentication "spin-fud/pkg/controllers"

	"github.com/gorilla/mux"
)

func GenerateRouter() {
	router := mux.NewRouter()
	authenticationRouters(router)
	listen(router, ":8000")
}

/*
pkg/controllers/authentication router.
*/
func authenticationRouters(router *mux.Router) {
	router.HandleFunc("/login", authentication.Authentication).Methods("POST")
	router.HandleFunc("/register", authentication.Register).Methods("POST")
}

func listen(router *mux.Router, port string) {
	http.ListenAndServe(port, router)
}
