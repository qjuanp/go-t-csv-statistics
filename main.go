package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This start here!")

	filePath := os.Args[1] // first argument will have the file path

	fmt.Printf("File path %s\n", filePath)

	people := readCsv(filePath)

	averageAge := calculateAverageAge(people)
	medianPerson := calculateMedian(people)

	fmt.Println(people.toString())
	fmt.Println(fmt.Sprintf("# People: %d", len(people)))
	fmt.Println(fmt.Sprintf("Average age: %d", averageAge))
	fmt.Println(fmt.Sprintf("Median: %s", medianPerson.toString()))
}

func calculateMedian(people People) Person {
	median := (len(people) + 1) / 2

	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	return people[median]
}

func calculateAverageAge(people People) int {
	sum := 0

	for _, person := range people {
		sum += person.age
	}

	return sum / len(people)
}

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

func asRecord(entry string) Person {
	fields := strings.SplitAfter(entry, ",")
	age, _ := strconv.Atoi(fields[2])
	return Person{fields[0], fields[1], age}
}
