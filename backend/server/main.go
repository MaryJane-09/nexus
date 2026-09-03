package main

import (
	"fmt"
	"net/http"
)

func healthHandelr(w http.ResponseWriter, r*http.Request) {
	fmt.Fprint(w, "Starting Nexus Server...")

}

func main() {
	fmt.Println("Nexus server running on :8080")
	http.HandleFunc("/health", healthHandelr)
	err := http.ListenAndServe(":8080", nil); if err != nil {
		fmt.Println("Failed to start HTTP server: ", err)
	}
}