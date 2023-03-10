package main

import (
	"fmt"

	"github.com/CarlosBarbosaFilho/api-rest-golan/database"
	"github.com/CarlosBarbosaFilho/api-rest-golan/routes"
)

func main() {
	fmt.Print("Server is running...")
	database.ConnectionDB()
	routes.HandleRequest()
}
