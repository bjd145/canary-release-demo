package main

import (
	"encoding/json"
	"time"
	"os"
	"net/http"
	"fmt"
	"log"
	"runtime"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var version string = "v1"

type OS struct {
	Time string
	AksHost string
	Host string
	Region string
	OSType string
	Version string
}

type newAPIHandler struct { }
func (eh *newAPIHandler) getOperatingSystemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	host, _ := os.Hostname()
	ostype := runtime.GOOS 
	
	aks_host := "UNKNOWN"
	if os.Getenv("AKSHOST") != "" {
		aks_host = os.Getenv("AKSHOST")
	}

	region := "UNKNOWN"
	if os.Getenv("REGION") != "" {
		region = os.Getenv("REGION")
	}

	msg := OS{ 
		time.Now().Format(time.RFC850), 
		aks_host,
		host,
		region,
		ostype,
		version}	

	json.NewEncoder(w).Encode(msg)
}

func (eh *newAPIHandler) optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {

	if os.Getenv("API_VERSION") != "" {
		version = os.Getenv("API_VERSION")
	}
	
	handler := newAPIHandler{}

	r := mux.NewRouter()
	apirouter := r.PathPrefix("/api").Subrouter()
	apirouter.Methods("GET").Path("/os").HandlerFunc(handler.getOperatingSystemHandler)
	apirouter.Methods("OPTIONS").Path("/os").HandlerFunc(handler.optionsHandler)

	server := cors.Default().Handler(r)

	port := ":8081"
	fmt.Print("Listening on port", port)
	log.Fatal(http.ListenAndServe( port , server))
}
