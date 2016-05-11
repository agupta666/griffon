package main

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"os"
	"fmt"
)

const BUCKET = "entries"

type Entry struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func (e Entry) Array() []string {
	return []string{ e.Name, e.IP, fmt.Sprintf("%d", e.Port) }
}

var (
	db *bolt.DB
)

func InitDB(path string, mode os.FileMode) (*bolt.DB, error) {
	var err error
	db, err = bolt.Open(path, mode, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func saveEntry(e *Entry) error {
	data, err := json.Marshal(e)

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BUCKET))
		if err != nil {
			return err
		}
		err = b.Put([]byte(e.Name), data)
		return err
	})

	return err
}

func allEntries() []*Entry {
	entries := make([]*Entry, 0)

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var entry Entry
			err := json.Unmarshal(v, &entry)
			if err != nil {
				return err
			}
			entries = append(entries, &entry)
		}

		return nil
	})

	return entries
}

func lookup(name string) (*Entry, error) {
	var entry Entry
	e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET))
		v := b.Get([]byte(name))

		err := json.Unmarshal(v, &entry)

		if err != nil {
			return err
		}

		return nil
	})

	return &entry, e
}
