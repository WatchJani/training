package main

import "fmt"

// bez pointera mogu samo da citam te podatke i kopiram ih
// sa pointerom mogu da ih mjenjam i citam i promjenice se na svakom mjestu a ne samo trenutnom
func RunTraining7() {
	one := 1
	two := 2

	store := map[string]*int{
		"1": &one,
		"2": &two,
	}

	for _, value := range store {
		*value = 210
		fmt.Println(value)
	}

	fmt.Println(store)
}
