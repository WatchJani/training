package main

import "fmt"

type City struct {
	name       string
	population int
}

func NewCity(name string, population int) *City {
	return &City{
		name:       name,
		population: population,
	}
}

func RunTraining2() {
	store := map[string]*City{
		"New York": NewCity("New York", 20_000),
	}

	for _, city := range store {
		city.population = 15_000
	}

	data := store["New York"]
	data.population = 15

	fmt.Println(store["New York"].population)
}
