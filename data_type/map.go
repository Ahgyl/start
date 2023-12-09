package data_type

import (
	"fmt"
)

func mapExample() {
	var firstMap map[string]bool
	firstMap = make(map[string]bool)
	firstMap["London"] = true
	firstMap["Paris"] = false
	firstMap["Tokio"] = true
	fmt.Println("Города где я была в хэш таблице", firstMap)
	for key, value := range firstMap {
		if value {
			fmt.Println("я была в городе", key)
		}
	}

	value, ok := firstMap["Kazan"]
	if !ok {
		fmt.Println("Казани нет в списке")
	} else {
		fmt.Println("Казань есть в списке и равно", value)
	}
	firstMap["Kazan"] = true
	value, ok = firstMap["Kazan"]
	if !ok {
		fmt.Println("Казани нет в списке до сих пор")
	} else {
		fmt.Println("Казань есть теперь в списке и равно", value)
	}
	fmt.Println("Количество определенных городов в хэш мапе равно", len(firstMap))
}
