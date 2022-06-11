package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type RefFile struct {
	Files []string `json:"files"`
}

func LocateUrls(refUrl url.URL) []url.URL {
	resp, httpErr := http.Get(refUrl.String())

	if httpErr != nil {
		log.Fatalf("Cannot get reference file from %s", refUrl.String())
	}

	defer resp.Body.Close()

	var ref RefFile
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatalf("Cannot read response %s", readErr)
	}

	fmt.Printf("Body response %s", string(body[:]))

	jsonErr := json.Unmarshal(body, &ref)

	if jsonErr != nil {
		log.Fatalf("Error marshalling %s", jsonErr)
	}

	fmt.Printf("Object with refs %+v\n", ref)

	var urls []url.URL

	for _, refUrl := range ref.Files {
		url, _ := url.Parse(refUrl)
		urls = append(urls, *url)
	}

	fmt.Printf("Urls: %v\n", urls)

	return urls
}
