package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getEntries(host string, port int) ([]Entry, error) {
	resp, err := http.Get(epUrl(host, port))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var entries []Entry
	err = json.Unmarshal(body, &entries)

	if err != nil {
		return nil, err
	}

	return entries, nil
}
