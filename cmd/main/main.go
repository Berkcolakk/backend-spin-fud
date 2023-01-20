package main

import (
	dbConnection "spin-fud/pkg/db"
	routes "spin-fud/pkg/routers"
)

func main() {
	routes.GenerateRouter()
	dbConnection.DBConnect()

}
