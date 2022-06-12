package main

import (
	"fmt"
	"strings"
)

func report(people People) string {
	rb := strings.Builder{}

	averageAge := avg(people)
	medianPerson := median(people)

	rb.WriteString(fmt.Sprintf("# People: %d\n", len(people)))
	rb.WriteString(fmt.Sprintf("Average age: %d\n", averageAge))
	rb.WriteString(fmt.Sprintf("Median: %s\n", medianPerson.toString()))

	return rb.String()
}
