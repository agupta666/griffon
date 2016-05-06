package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func listEntries(w http.ResponseWriter, r *http.Request) {
	entries := allEntries()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entries)
}

func addEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e Entry
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &e)

	saveEntry(&e)

	json.NewEncoder(w).Encode(e)
}

func StartRESTServer(host string, port int) {
	router := mux.NewRouter()
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/entries", listEntries).Methods("GET")
	apiV1.HandleFunc("/entries", addEntry).Methods("POST")
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(fmt.Sprintf("%s:%d", host, port))
}
