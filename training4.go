package main

import "fmt"

func RunTraining4() {
	store := map[string]*[]int{
		"1": {
			1, 2, 4, 5, 47,
		},
	}

	data := store["1"]

	*data = []int{}

	fmt.Println(data)
}
