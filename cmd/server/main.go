package main

import (
	"log"

	"github.com/JonathanBerhe/distributedlog/internal/server"
)

func main() {
	server := server.NewHTTPServer(":8080")
	log.Fatal(server.ListenAndServe())
}
