package main

import "sort"

func median(people People) Person {
	median := (len(people) + 1) / 2

	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	return people[median]
}

func avg(people People) int {
	sum := 0

	for _, person := range people {
		sum += person.age
	}

	return sum / len(people)
}
