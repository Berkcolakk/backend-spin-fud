package main

import (
	dbConnection "spin-fud/pkg/db"
	// routes "rest-api-test/pkg/routers"
)

func main() {
	// routes.GenerateRouter()
	dbConnection.DBConnect()

}
