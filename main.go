package main

import (
	"flag"
	"log"
	"os"

	"github.com/miekg/dns"
)

var (
	domain      = flag.String("r", "service.consul.", "root domain.")
	dnsHost     = flag.String("b", "0.0.0.0", "bind address.")
	dnsPort     = flag.Int("p", 8053, "port to listen.")
	restHost    = flag.String("s", "0.0.0.0", "bind address for REST interface.")
	restPort    = flag.Int("q", 8080, "port to listen for REST interface.")
	interactive = flag.Bool("c", false, "start in interactive mode.")
)

func main() {

	flag.Parse()

	if *interactive {
		f, err := os.OpenFile("griffon.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("ERROR: opening log file", err)
			return
		}
		defer f.Close()

		log.SetOutput(f)
	}

	db, err := InitDB(".griffon.db", 0600)

	if err != nil {
		log.Println("ERROR: opening database", err)
		return
	}
	defer db.Close()

	go StartRESTServer(*restHost, *restPort)

	dns.HandleFunc(*domain, Handler)
	go serveDNS("tcp", *dnsHost, *dnsPort)

	if *interactive {
		go serveDNS("udp", *dnsHost, *dnsPort)
		startShell()
	} else {
		serveDNS("udp", *dnsHost, *dnsPort)
	}

}
