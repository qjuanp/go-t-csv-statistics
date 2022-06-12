package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func readRemoteFile(url url.URL) io.Reader {
	resp, httpErr := http.Get(url.String())

	logIfError(httpErr, fmt.Sprintf("Error requesting the file from %s", url.String()))

	return resp.Body
}

func readLocalFile(path string) io.Reader {
	file, err := os.Open(path)

	logIfError(err, fmt.Sprintf("Error reading the file from %s", path))

	return file
}

func logIfError(err error, message string) {
	if err != nil {
		log.Fatalf("%s\nError:%s", message, err)
	}
}
