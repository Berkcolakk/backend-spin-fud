package authentication

import (
	"fmt"
	"net/http"
)

func Authentication(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Test")
}

func Register(w http.ResponseWriter, req *http.Request) {

}
