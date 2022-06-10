package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

type People []Person

func main() {
	fmt.Println("This start here!")

	filePath := os.Args[1] // first argument will have the file path

	fmt.Printf("File path %s\n", filePath)

	people := readCsv(filePath)

	fmt.Println(people.toString())
}

func readCsv(path string) People {
	var people People
	file, err := os.Open(path)

	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		people = append(people, asRecord(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return people
}

func asRecord(entry string) Person {
	fields := strings.SplitAfter(entry, ",")
	age, _ := strconv.Atoi(fields[2])
	return Person{fields[0], fields[1], age}
}

func (people People) toString() string {
	builder := strings.Builder{}

	for _, person := range people {
		builder.WriteString(fmt.Sprintf("%s\n", person.toString()))
	}

	return builder.String()
}

func (person Person) toString() string {
	return fmt.Sprintf("Person -> name: %s | lastname: %s | age: %d", person.firstName, person.lastName, person.age)
}
