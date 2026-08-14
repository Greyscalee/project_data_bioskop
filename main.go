package main

import (
	"data-bioskop/config"
	"data-bioskop/routers"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
