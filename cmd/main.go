package main

import (
	"client-server-api/client"
	"client-server-api/server"
)

func main() {
	go server.NewServer()

	client.BuscarCotacao()
}
