package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This start here!")

	filePath := os.Args[1] // first argument will have the file path

	fmt.Printf("File path %s\n", filePath)

	readCsv(filePath)
}

func readCsv(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Print(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
