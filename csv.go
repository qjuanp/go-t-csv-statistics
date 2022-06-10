package main

import (
	"bufio"
	"fmt"
	"os"
)

func readCsv(path string) People {
	var people People
	file, err := os.Open(path)

	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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
