package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func epUrl(host string, port int) string {
	return fmt.Sprintf("http://%s:%d/api/v1/entries", host, port)
}

func writer(name string, format string) (*os.File, error) {
	fname := fmt.Sprintf("%s.%s", name, format)
	f, err := os.Create(fname)

	return f, err
}

func writeJson(w io.Writer, entries []Entry) error {
	jsData, err := json.MarshalIndent(entries, "", "\t")

	if err != nil {
		return err
	}
	w.Write(jsData)
	return nil
}

func writeCSV(w io.Writer, entries []Entry) error {

	csvw := csv.NewWriter(w)

	for _, e := range entries {
		err := csvw.Write(e.Array())

		if err != nil {
			return err
		}
	}
	csvw.Flush()
	if err := csvw.Error(); err != nil {
		return err
	}

	return nil
}

func exportData(c *cli.Context) {

	format := c.String("f")
	fname := c.String("n")
	restHost := c.String("s")
	restPort := c.Int("p")

	resp, err := http.Get(epUrl(restHost, restPort))

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: {HTTP GET}", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:", err)
		return
	}

	var entries []Entry
	err = json.Unmarshal(body, &entries)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:", err)
		return
	}

	w, err := writer(fname, format)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:", err)
		return
	}

	switch format {
	case "json":
		err = writeJson(w, entries)
	case "csv":
		err = writeCSV(w, entries)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}
}
