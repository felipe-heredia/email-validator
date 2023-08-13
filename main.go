package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Printf("Write the domain to check: ")

	var domain string

	if _, err := fmt.Scanln(&domain); err != nil && !strings.Contains(err.Error(), "unexpected newline") {
		log.Fatal(err)
	}

	if domain == "" {
		log.Fatal("Please, inform a domain")
	}

	fmt.Print("Verifying domain... \n")
	checkDomain(domain)
}

func checkDomain(domain string) {
	var hasMx, hasSPF bool
	var spfRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	fmt.Printf("domain: %v \n", domain)
	fmt.Printf("hasMx: %v \n", hasMx)
	fmt.Printf("hasSpf: %v \n", hasSPF)
	fmt.Printf("spfRecord: %v \n", spfRecord)
}
