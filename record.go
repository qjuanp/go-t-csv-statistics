package main

import (
	"fmt"
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
