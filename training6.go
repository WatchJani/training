package main

import "fmt"

func RunTraining6() {
	data := &[]int{1, 2, 3, 4, 5}

	//POINTER NA SLICE NE NA VALUE
	for _, value := range *data {
		value = 1
		fmt.Println(value)
	}

	(*data)[0] = 3

	fmt.Println(data)

	//============================
}
