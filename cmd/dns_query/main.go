package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

const (
	Timeout = 5
	Port    = 53
)

type TimedResult interface {
	fmt.Stringer
	Duration() time.Duration
}

type result struct {
	ip  string
	rtt time.Duration
}

func (r *result) String() string {
	return r.ip
}

func (r *result) Duration() time.Duration {
	return r.rtt
}

var (
	domain, server string
)

func init() {
	flag.StringVar(&domain, "domain", "", "domain name")
	flag.StringVar(&server, "server", "", "server name")
}

func main() {
	flag.Parse()
	if domain == "" || server == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	// fmt.Printf("domain %s, server %s", domain, server)
	res, err := getDnsQuery(domain, server)
	if err != nil {
		log.Print(err)
		os.Exit(1)
		return
	}
	fmt.Printf("dns_query_res,domain=%s,record_type=A,server=%s,ip=%s query_time_ms=%.6f\n",
		domain, server, res.String(), float64(res.Duration())/1e6)
}

func getDnsQuery(domain string, server string) (TimedResult, error) {

	c := new(dns.Client)
	c.ReadTimeout = time.Duration(Timeout) * time.Second

	m := new(dns.Msg)

	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	m.RecursionDesired = true

	r, rtt, err := c.Exchange(m, net.JoinHostPort(server, strconv.Itoa(Port)))
	if err != nil {
		return nil, err
	}
	if r.Rcode != dns.RcodeSuccess {
		return nil, fmt.Errorf("Invalid answer name %s after %v query for %s\n", domain, dns.TypeA, domain)
	}

	for _, a := range r.Answer {
		if h, ok := a.(*dns.A); ok {
			return &result{h.A.String(), rtt}, nil
		}
	}

	return nil, fmt.Errorf("got invalid result: %v", r.Answer)

}
