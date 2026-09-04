package main

import (
	"fmt"
	"net/http"
	"github.com/MaryJane-09/nexus/backend/config"
	"github.com/MaryJane-09/nexus/backend/internal/router"
)

func main() {
	fmt.Println("Nexus server running on", config.ServerPort)
	server := router.New()
	err := http.ListenAndServe(config.ServerPort, server)
	if err != nil {
		fmt.Println("Failed to start HTTP server: ", err)
	}
}
