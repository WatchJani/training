package main

import "fmt"

func RunTraining5() {
	dubica := NewCity("Koz. Dubica", 18_000)

	//pointer na strukturu je super kada zelimo da se ta vrijednost objekta
	//promjeni na sve i jednom mjestu odjednom ili da imas vrijednost nil,
	// jer kada ne bi imali poinetr
	//vrijednost bi se promjenila samo u mapi
	city := map[string][]*City{
		"1": {
			dubica,
		},
	}

	p := city["1"]

	p[0].name = "top"

	fmt.Print(dubica)

	//================================================================================================
	grad := map[string]*[]City{
		"1": {
			*dubica,
		},
	}

	c := grad["1"]

	(*c)[0].name = "free"

	fmt.Print(dubica)
	fmt.Println(c)
}
