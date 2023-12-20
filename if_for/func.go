package if_for

import (
	"fmt"
)

func ExampleFunc() {
	sumResult := sum(1, -9)
	fmt.Println("Результат сложения двух числе 1 и -9 функцией sum", sumResult)

	//Ссылка на функцию Woof
	dog := Woof
	//Ссылка на функцию Miay
	cat := Miay

	//Т.к. dog это ссылка на функцию Woof, a cat на функцию Miya, а Child это функция,
	//которая принимает функции типа Sayers, а Woof и Miya функции типа Sayers
	Child(dog)
	Child(cat)
}

// Неименованный возвращаемый результат
func sum(a int, b int) int {
	result := a + b
	return result
}

// Именованный возвращаемый результат
func sum2(a int, b int) (res int) {
	res = a + b
	return
}

// Тип функции, которые ничего не принимают и возвращают строку
type Sayers func() string

func Woof() string {
	return "woof woof rrrr"
}

func Miay() string {
	return "miay miay shhh"
}

// Функция принимает функцию типа sayers и уже внутри ее вызывает
func Child(animal Sayers) {
	fmt.Println("Mommy, i got animal, its says", animal())
}
