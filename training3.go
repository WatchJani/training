package main

import (
	"fmt"
)

func RunTraining3() {
	store := map[string]*[]int{
		"1": {
			1, 2, 4, 5, 47,
		},
	}

	for _, data := range store {
		*data = []int{}
		fmt.Println(data)
	}

	fmt.Println(store["1"])
}
