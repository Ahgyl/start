package data_type

import (
	"fmt"
)

func Example() {
	// С плавающей точкой
	var f float32 = 18
	var g float32 = 4.5
	var d float64 = 0.23
	var pi float64 = 3.14
	var e float64 = 2.7
	fmt.Println("Примеры чисел с плавающей точкой", f, g, d, pi, e)

	// Логические
	var isAlive bool = true
	var isEnabled = false
	isDied := true
	fmt.Println("Примеры логических переменных", isAlive, isEnabled, isDied)

	// Строки
	var name string = "Том Сойер"
	fmt.Println("Переменная с типом строка:", name)
	color := "blue"
	fmt.Println("Переменная строка с кратким объявлением:", color)

	// Сложение (int, float и string)
	var a float64 = 8.1
	b := 8.0
	c := a + b
	fmt.Println(fmt.Sprintf("Результат сложения %.2f и %.3f равен %f", a, b, c))

	name = name + " любимая книга"
	fmt.Println("Строка после конкатинации:", name)

	// Вычитание (int и float)

	// Деление (int и float)
	// ТАК НЕЛЬЗЯ! x должен быть дробным числом
	// var x int
	// var y, z float64
	// x := y / z

	// Умножение (int и float)

	// ТЕМА: логические операции
	var bl bool
	bl = 5 == 5
	fmt.Println("Результат определения равенства 5 и 5", bl)
	val1 := 5
	val2 := 6
	bl = val1 != val2
	fmt.Println(fmt.Sprintf("%d не равно %d = %t", val1, val2, bl))
}

// Целочисленные
func integerExample() {
	// 1 Объявление переменной через var
	var myValue int64
	myValue = 900
	fmt.Println("Объявление int64 переменной через var", myValue)
	// 2 Сокращенное объявление целочисленной переменной через :=
	myValue2 := 9
	fmt.Println("Сокращенное объявление целочисленной переменной через :=", myValue2)
	// 3 Объявление переменной через var с последующим присвоением значения
	var myValue3 int32 = 100
	fmt.Println("Объявление переменной через var с последующим присвоением значения", myValue3)
	// Присвоение нового значения переменной
	myValue3 = myValue3 + 100
	fmt.Println("Присвоение нового значения переменной", myValue3)
	// Создание двух переменных
	var (
		aa int = 1
		bb int = 2
	)
	fmt.Println("Создание двух переменных через var()", aa, bb)
}
