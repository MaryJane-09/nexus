package health

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r*http.Request) {
	fmt.Fprint(w, "Starting Nexus Server...")

}