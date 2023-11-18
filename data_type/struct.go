package data_type

import (
	"fmt"
)

type Person struct {
	// Имя человека
	Name   string
	Age    int
	weight float64
}

func StructExample() {
	ahgyl := Person{
		Name:   "Ahgyl",
		Age:    18,
		weight: 60,
	}
	rafael := Person{
		Name:   "Rafael",
		Age:    31,
		weight: 74,
	}
	fmt.Println(fmt.Sprintf("Ahgyl=%v; Rafael=%v", ahgyl, rafael))

	persons := []Person{ahgyl}
	persons = append(persons, rafael)

	shamil := Person{"Shamil", 0, 7.885}
	persons = append(persons, shamil)

	personsMap := make(map[string]Person)
	for _, value := range persons {
		personsMap[value.Name] = value
	}
	fmt.Println(fmt.Sprintf("Наша полученная мапа =%v", personsMap))
}
