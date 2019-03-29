package main

import (
	"flag"
	"log"
	"strings"

	"lucas/crawl"
	"lucas/validate"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Please provide a single url value.")
	}

	address := args[0]
	if !validate.URLAddress(address) {
		log.Fatal("Please provide a valid url value.")
	}

	page, err := crawl.Download(address)
	if err != nil {
		log.Fatalf("Failed to download page: %s", err.Error())
	}

	for _, l := range strings.Split(page, "\n") {
		u, ok := crawl.ExtractURL(l)
		if ok {
			println(u)
		}
	}
}
