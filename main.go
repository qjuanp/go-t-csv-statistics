package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("This start here!")

	filePath := os.Args[1] // first argument will have the file path

	fmt.Printf("File path %s\n", filePath)

	people := readCsv(filePath)

	averageAge := avg(people)
	medianPerson := median(people)

	fmt.Println(people.toString())
	fmt.Println(fmt.Sprintf("# People: %d", len(people)))
	fmt.Println(fmt.Sprintf("Average age: %d", averageAge))
	fmt.Println(fmt.Sprintf("Median: %s", medianPerson.toString()))
}
