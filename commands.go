package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func listData(c *cli.Context) {
	restHost := c.String("s")
	restPort := c.Int("p")

	entries, err := getEntries(restHost, restPort)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: fetching entries", err)
		return
	}

	for _, e := range entries {
		fmt.Println(e)
	}
}

func addData(c *cli.Context) {
	restHost := c.String("s")
	restPort := c.Int("p")

	name := c.String("name")
	ip := c.String("ip")
	port := c.Int("port")

	entry := Entry{Name: name, IP: ip, Port: port}

	err := addEntry(restHost, restPort, entry)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: fetching entries", err)
	} else {
		fmt.Fprintln(os.Stdout, "added entry", entry)
	}

}
