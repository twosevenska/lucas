package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please provide a single url value.")
		os.Exit(1)
	}

	resp, err := http.Get(args[0])
	if err != nil {
		fmt.Errorf("Failed to get webpage: %s", err.Error())
	}
	defer resp.Body.Close()

	// Warning: This will not handle large webpages properly
	// https://groups.google.com/forum/#!topic/golang-nuts/sAwDldpkMGQ
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Failed to read webpage: %s", err.Error())
	}

	fmt.Println(string(html))

	//TODO: Call our crawl package here
}
