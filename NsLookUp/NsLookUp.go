package NsLookUp

import (
	"fmt"
	"net"
)

func LookUp(domain string) {
	nameserver, _ := net.LookupNS(domain)
	for _, ns := range nameserver {
		fmt.Println(ns)
	}
	fmt.Println("LoopUpIP ...........")
	iprecords, _ := net.LookupIP(domain)
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	fmt.Println("LoopUpCNAME ...........")
	cnames, _ := net.LookupCNAME(domain)
	fmt.Println(cnames)
}
