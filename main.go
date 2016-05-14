package main

import (
	"github.com/codegangsta/cli"
	"github.com/miekg/dns"
	"log"
	"os"
)

/*
var (
	root = flag.String("r", "service.consul.", "root domain.")
	host = flag.String("b", "0.0.0.0", "bind address.")
	port = flag.Int("p", 8053, "port to listen.")
)
*/
func serve(c *cli.Context) {

	domain := c.String("d")
	dnsHost := c.String("b")
	dnsPort := c.Int("p")
	restHost := c.String("s")
	restPort := c.Int("q")

	db, err := InitDB(".griffon.db", 0600)

	if err != nil {
		log.Println("ERROR: opening database", err)
		return
	}
	defer db.Close()

	go StartRESTServer(restHost, restPort)

	dns.HandleFunc(domain, Handler)
	go serveDNS("tcp", dnsHost, dnsPort)
	serveDNS("udp", dnsHost, dnsPort)
}

func main() {
	app := cli.NewApp()
	app.Name = "griffon"
	app.Usage = "dns server"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "start dns server",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "d", Value: "service.consul", Usage: "domain."},
				cli.StringFlag{Name: "b", Value: "0.0.0.0", Usage: "bind address for dns server."},
				cli.IntFlag{Name: "p", Value: 8053, Usage: "dns server port."},
				cli.StringFlag{Name: "s", Value: "0.0.0.0", Usage: "bind address for REST server."},
				cli.IntFlag{Name: "q", Value: 3000, Usage: "REST server port."},
			},
			Action: serve,
		},
		{
			Name:   "import",
			Usage:  "import entries from csv / json.",
			Action: importData,
		},
		{
			Name:   "list",
			Usage:  "print all entries.",
			Action: listData,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "s", Value: "0.0.0.0", Usage: "address of the REST server."},
				cli.IntFlag{Name: "p", Value: 3000, Usage: "REST server port."},
			},
		},
		{
			Name:   "add",
			Usage:  "add entry.",
			Action: addData,
		},
		{
			Name:  "export",
			Usage: "export data to csv / json.",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "f", Value: "json", Usage: "export entries to given format. [csv|json]"},
				cli.StringFlag{Name: "n", Value: "data", Usage: "name of the output file."},
				cli.StringFlag{Name: "s", Value: "0.0.0.0", Usage: "address of the REST server."},
				cli.IntFlag{Name: "p", Value: 3000, Usage: "REST server port."},
			},
			Action: exportData,
		},
	}

	app.Run(os.Args)

}
