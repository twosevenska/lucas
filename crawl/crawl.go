package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Download fetches a webpage and returns it as a string
func Download(address string) (string, error) {
	resp, err := http.Get(address)
	if err != nil {
		return "", fmt.Errorf("failed to get webpage: %s", err.Error())
	}
	defer resp.Body.Close()

	// Warning: This will not handle large webpages properly
	// https://groups.google.com/forum/#!topic/golang-nuts/sAwDldpkMGQ
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read webpage: %s", err.Error())

	}

	return string(html), nil
}

// ExtractURL attemps to extract a find a url tag and extract the address
func ExtractURL(line string) (string, bool) {
	startIndex := strings.Index(line, "a href=\"") + 8
	if startIndex == -1 {
		return "", false
	}

	endIndex := strings.LastIndex(line, "\">")
	if endIndex == -1 {
		return "", false
	}

	result := line[startIndex:endIndex]
	if result == "" {
		return "", false
	}
	return result, true
}
