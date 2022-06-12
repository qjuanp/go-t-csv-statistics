package main

import (
	"net/url"
	"testing"
)

func TestFileLocateUrls(t *testing.T) {
	expectedUrl, _ := url.Parse("https://qjuanp.dev/static-data-omelette/p-age/data/10-records.csv")
	refUrl, _ := url.Parse("https://qjuanp.dev/static-data-omelette/p-age/ref/one-file-10p.json")

	urls := LocateUrls(*refUrl)

	if len(urls) != 1 {
		t.Errorf("Expected just one url, resolved: %d", len(urls))
	}

	if urls[0].String() != expectedUrl.String() {
		t.Errorf("Expected url don't match\nr:%s\ne:%s", urls[0].String(), expectedUrl.String())
	}
}
