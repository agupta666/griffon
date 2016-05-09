package main

import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"log"
)

func Handler(w dns.ResponseWriter, r *dns.Msg) {
	log.Println(r)

	e, err := lookup(r.Question[0].Name)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	m := new(dns.Msg)
	m.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		s := fmt.Sprintf("%s  0 IN    A    %s", e.Name, e.IP)
		rr, _ := dns.NewRR(s)
		m.Answer = append(m.Answer, rr)
	case dns.TypeSRV:
		s := fmt.Sprintf("%s    0    IN    SRV    1 1 %d %s", e.Name, e.Port, e.Name)
		sa := fmt.Sprintf("%s 0 IN    A    %s", e.Name, e.IP)
		rr, _ := dns.NewRR(s)
		rr1, _ := dns.NewRR(sa)
		m.Answer = append(m.Answer, rr)
		m.Extra = append(m.Extra, rr1)
	}
	w.WriteMsg(m)
}

func serve(net string, host string, port int) {
	log.Printf("starting %s server on %s:%d", net, host, port)
	server := &dns.Server{Addr: fmt.Sprintf("%s:%d", host, port), Net: net, TsigSecret: nil}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to setup the "+net+" server: %s\n", err.Error())
	}
}

var (
	root = flag.String("r", "service.consul.", "root domain.")
	host = flag.String("b", "0.0.0.0", "bind address.")
	port = flag.Int("p", 8053, "port to listen.")
)

func main() {
	flag.Parse()
	go StartRESTServer("0.0.0.0", 3000)

	dns.HandleFunc(*root, Handler)
	go serve("tcp", *host, *port)
	serve("udp", *host, *port)
}
