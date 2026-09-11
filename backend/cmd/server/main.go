package main

import (
	"fmt"
	"net/http"

	"github.com/MaryJane-09/nexus/backend/config"
	"github.com/MaryJane-09/nexus/backend/internal/router"
)

func main() {

	fmt.Println("Nexus server running on port", config.AppConfig.ServerPort)
	fmt.Println("Gmail address loaded:", config.AppConfig.GmailAddress != "")
	fmt.Println("App password loaded:", config.AppConfig.GmailAppPassword != "")
	server := router.New()
	err := http.ListenAndServe(config.AppConfig.ServerPort, server)
	if err != nil {
		fmt.Println("Failed to start HTTP server: ", err)
	}
}
