package main

import "fmt"

func RunTraining8() {
	slice := &[]int{1, 2, 3}

	store := map[string]*[]int{
		"1": slice,
	}

	data := store["1"]

	*data = append(*data, 1)

	fmt.Println(slice)
}
