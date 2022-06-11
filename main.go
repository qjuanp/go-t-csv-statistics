package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Json struct {
		RefUrl string `arg name:"refurl" help:"Define json that references the csv files to process"`
	} `cmd help:"Process referenced csv files in a json"`
	Csv struct {
		Paths []string `arg name:"paths" help:"Pass each csv file to process"`
	} `cmd help:"Process individual csv files"`
}

func main() {
	ctx := kong.Parse(&CLI)
	var people People
	switch ctx.Command() {
	case "json <refurl>":
		fmt.Printf("Use reference request to: %s", CLI.Json.RefUrl)
		refUrl, _ := url.Parse(CLI.Json.RefUrl)
		csvUrls := LocateUrls(*refUrl)
		for _, csvUrl := range csvUrls {
			people = append(people, mapCsv(readRemoteFile(csvUrl))...)
		}
	case "csv <paths>":
		fmt.Printf("Read %d csv files.\nPaths: %v", len(CLI.Csv.Paths), CLI.Csv.Paths)
		for _, csvPath := range CLI.Csv.Paths {
			if _, err := os.Stat(csvPath); !os.IsNotExist(err) { // File can be accessed
				people = append(people, mapCsv(readLocalFile(csvPath))...)
			} else if url, err := url.Parse(csvPath); err == nil { // Url can be parsed
				people = append(people, mapCsv(readRemoteFile(*url))...)
			} else {
				log.Fatalf("Parameter not recognizes %s", csvPath)
			}
		}
	default:
		panic(ctx.Command())
	}
	fmt.Println(report(people))
}
