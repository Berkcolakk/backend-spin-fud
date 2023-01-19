package main

import (
	dbConnection "rest-api-test/pkg/db"
	// routes "rest-api-test/pkg/routers"
)

func main() {
	// routes.GenerateRouter()
	dbConnection.DBConnect()

}
