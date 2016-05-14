package main

import (
	"github.com/codegangsta/cli"
	"os"
	"fmt"
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
