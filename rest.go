package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func listEntriesHandler(w http.ResponseWriter, r *http.Request) {
	entries := allEntries()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entries)
}

func addEntryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e Entry
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &e)
	log.Println("entry", e)

	err := saveEntry(&e)

	if err != nil {
		log.Println("ERROR: saving entry.", err)
	}

	json.NewEncoder(w).Encode(e)
}

func StartRESTServer(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)

	router := mux.NewRouter()
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/entries", listEntriesHandler).Methods("GET")
	apiV1.HandleFunc("/entries", addEntryHandler).Methods("POST")
	err := http.ListenAndServe(addr, router)

	if err != nil {
		log.Println("ERROR: starting REST interface", err)
	}
}
