package main

import (
	"flag"
	"fmt"

	"lucas/crawl"
	"lucas/validate"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please provide a single url value.")
		return
	}

	address := args[0]
	if !validate.URLAddress(address) {
		fmt.Println("Please provide a valid url value.")
		return
	}

	fmt.Println(crawl.Download(address))
}
