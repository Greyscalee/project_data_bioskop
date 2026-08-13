package main

import "data-bioskop/routers"

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}

