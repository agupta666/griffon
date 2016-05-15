package main

import (
	"bytes"
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

func addEntry(host string, port int, entry Entry) error {
	var payload, err = json.Marshal(entry)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", epUrl(host, port), bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	return nil
}
