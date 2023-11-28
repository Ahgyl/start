package if_for

import (
	"fmt"
)

// Условие ЕСЛИ и ИНАЧЕ
func ExampleIf(val, val2 int) {
	if val > 0 {
		fmt.Println(fmt.Sprintf("Val %d Больше нуля", val))
	} else {
		fmt.Println(fmt.Sprintf("Val %d Меньше нуля", val))
	}
	//Можно задать if без else в случае когда нужно выполнить только при выполнении условия true
	if val2 > 0 {
		fmt.Println(fmt.Sprintf("Val2 %d Больше нуля", val2))
	}
}

// Циклы
func ExampleFor(slice []int) {
	//Цикл по слайсу с итератором
	for i := 0; i < len(slice); i = i + 1 {
		fmt.Print("i =", i)
		fmt.Println(" slice[i] =", slice[i])
	}
	//Цикл по всему слайсу с range (нельзя задать границы)
	for i, v := range slice {
		fmt.Print("i =", i)
		fmt.Println(" v =", v)
	}
}
