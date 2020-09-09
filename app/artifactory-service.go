package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sonasingh46/artifactory-service/pkg/client"
	"io"
	"log"
	"net/http"
)

func main() {
	StartService()
}

// StartService starts a new artifactory proxy service
func StartService() {
	r := mux.NewRouter()
	// Endpoint for checking service liveness
	r.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")

	// Endpoint for listing all the albums.
	r.HandleFunc("/getLeastDownloaded", getLeastDownloaded).Methods("GET")
	log.Println("Setting up Jfrog Artifact client...")
	err := client.SetArtifactoryClient()
	if err != nil {
		log.Fatalf("failed to set Jfrog artifactory client:{%s}", err.Error())
	}
	log.Println("Artifact proxy running...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// HealthCheckHandler is the health check handler.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

// getLeastDownloaded returns the least 2 downloaded artifactory
// Note : The AQL for fetching the result is hardcoded and can be improved
// More info here : pkg/aql/aql.go
func getLeastDownloaded(w http.ResponseWriter, r *http.Request) {
	topDownloaded, err := client.GetArtifactoryClient().GetLeastDownloaded()
	if err != nil {
		log.Print("failed to get top downloaded artifacts", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to get top downloaded artifacts:"+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(topDownloaded)
	if err != nil {
		log.Print("failed to encode top downloaded artifacts", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to encode top downloaded artifact:"+err.Error())
		return
	}
}
