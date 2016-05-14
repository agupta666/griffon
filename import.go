package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func importJSON(fname string) {
	file, err := os.Open(fname)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}

	var entries []Entry
	json.Unmarshal(b, &entries)

	for _, e := range entries {
		var payload, err = json.Marshal(e)

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			return
		}

		req, err := http.NewRequest("POST", epUrl("0.0.0.0", 3000), bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			return
		}
		defer resp.Body.Close()
	}

}

func importCSV(fname string) {

}

func importData(c *cli.Context) {
	for _, fname := range c.Args() {
		format := filepath.Ext(fname)
		switch format {
		case ".json":
			importJSON(fname)
		case ".csv":
			importCSV(fname)
		default:
			fmt.Fprintf(os.Stderr, "ERROR: unsupported format %s.\n", format)
		}
	}
}
