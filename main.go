package main

import (
	"fmt"
	"net/url"

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
		fmt.Printf("Json processin %s [%+v]", ctx.Command(), CLI)
		refUrl, _ := url.Parse(CLI.Json.RefUrl)
		csvUrls := LocateUrls(*refUrl)
		for _, csvUrl := range csvUrls {
			people = append(people, mapCsv(readRemoteFile(csvUrl))...)
		}
		fmt.Println(report(people))
	case "csv <paths>":
		fmt.Printf("csv processin %s [%+v]", ctx.Command(), CLI)
	default:
		panic(ctx.Command())
	}
}
