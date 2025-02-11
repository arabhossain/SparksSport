package main

import (
	"SparksSport/cmd/app/bootstrap"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Initialize app and auto-register services
	container, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}

	// Start the server
	serverPort := strconv.Itoa(container.Config.ServerPort)
	fmt.Println("Server running on port: " + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, container.Router))
}
