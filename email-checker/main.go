package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal("[check-domain] err:", err.Error())
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	sprRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal("[check-domain] err:", err.Error())
	}

	for _, rec := range sprRecords {
		if strings.HasPrefix(rec, "v=spf1") {
			hasSPF = true
			sprRecord = rec
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatal("[check-domain] err:", err.Error())
	}

	for _, rec := range dmarcRecords {
		if strings.HasPrefix(rec, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = rec
			break
		}
	}

	fmt.Println("Domain:", domain)
	fmt.Println("hasMX:", hasMX)
	fmt.Println("haSPF:", hasSPF)
	fmt.Println("sprRecord:", sprRecord)
	fmt.Println("hasDMARK:", hasDMARC)
	fmt.Println("dmarkRecord:", dmarcRecord)
}

func main() {

	fmt.Println("----------- email-checker start ---------")

	fmt.Print("Enter domain (eg. gmail.com):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp := scanner.Text()

		checkDomain(inp)

		err := scanner.Err()
		if err != nil {
			log.Fatal("[scanner] err: ", err.Error())
		}
		fmt.Print("Enter domain (eg. gmail.com):")

	}

	fmt.Println("----------- email-checker end -----------")
}
