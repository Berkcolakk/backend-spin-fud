package authentication

import (
	"fmt"
	"net/http"
)

func Authentication(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Authenication")
}

func Register(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Register")
}
