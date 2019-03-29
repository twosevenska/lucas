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
	// It might be better to temporarily download to a file and scan it
	// line by line
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read webpage: %s", err.Error())

	}

	return string(html), nil
}

// extractURL attemps to extract a find a url tag and extract the address
func extractURL(line string) (string, bool) {
	startIndex := strings.Index(strings.ToLower(line), "href=\"")
	if startIndex == -1 {
		return "", false
	}

	trimmedPrefix := line[startIndex+6:]

	endIndex := strings.Index(trimmedPrefix, "\"")
	if endIndex == -1 {
		return "", false
	}

	result := trimmedPrefix[:endIndex]
	if result == "" {
		return "", false
	}
	return result, true
}

// ExtractURLS finds and extracts urls in page from hyperlink tags
func ExtractURLS(page string) map[string]int {
	lines := make(chan string)
	defer close(lines)

	done := make(chan bool)
	defer close(done)

	addresses := make(map[string]int)
	go func() {
		for {
			select {
			case l := <-lines:
				u, extracted := extractURL(l)
				if extracted {
					addresses[u] = addresses[u] + 1
				}
			case d := <-done:
				if d {
					return
				}
			}
		}
	}()

	for _, l := range strings.Split(page, "\n") {
		lines <- l
	}
	done <- true

	return addresses
}
