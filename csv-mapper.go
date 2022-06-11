package main

import (
	"bufio"
	"fmt"
	"io"
)

func mapCsv(reader io.Reader) People {
	var people People
	scanner := bufio.NewScanner(reader)

	for i := 0; scanner.Scan(); i++ {
		if i != 0 {
			people = append(people, asRecord(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return people
}
