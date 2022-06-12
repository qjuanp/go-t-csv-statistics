package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

type People []Person

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

func asRecord(entry string) Person {
	fields := strings.SplitAfter(entry, ",")
	age, _ := strconv.Atoi(fields[2])
	return Person{fields[0], fields[1], age}
}
