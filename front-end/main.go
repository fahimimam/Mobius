package main

import (
	"frontend/server"
)

func main() {
	router := GetRouter()
	server.HandleServer(router)
}
