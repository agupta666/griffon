package main

import (
	"errors"
	"sync"
)

type Entry struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

var (
	entries = make(map[string]*Entry)
	guard   sync.Mutex
)

func saveEntry(e *Entry) {
	guard.Lock()
	entries[e.Name] = e
	guard.Unlock()
}

func allEntries() []*Entry {
	exs := make([]*Entry, 0)
	for _, value := range entries {
		exs = append(exs, value)
	}
	return exs
}

func lookup(name string) (*Entry, error) {

	entry, ok := entries[name]

	if ok {
		return entry, nil
	} else {
		return nil, errors.New("lookup failed")
	}
}
