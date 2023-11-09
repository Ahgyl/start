package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	iLoveShamil bool = true
	age         int  = 18
	size             = 22
)

func main() {
	integerExample()
	/*
		// С плавающей точкой
		// var f float32 = 18
		// var g float32 = 4.5
		// var d float64 = 0.23
		// var pi float64 = 3.14
		// var e float64 = 2.7

		// Логические
		// var isAlive bool = true //
		// isAlive = false
		// var isEnabled bool = false

		// Строки
		var name string = "Том Сойер"
		name = name + " любимая книга"
		fmt.Println(name)
		//color := "blue"

		// +
		// int
		// float
		var a float64 = 8.1
		b := 8.0
		c := a + b
		fmt.Println(c)
		// string

		// -
		// int
		// float

		// /
		// int
		// float
		// todo ТАК НЕЛЬЗЯ! x должен быть дробным числом
		// var x int
		// var y, z float64
		// x := y / z

		// *
		// int
		// float

		// логические операции
		// bool

		var f float32
		fmt.Println("my empty float32 =", f)

		// ТЕМА: логические операции
		var bl bool
		bl = 5 == 5
		fmt.Println(bl)
		val1 := 5
		val2 := 6
		bl = val1 != val2
		fmt.Println(bl)

		// ТЕМА: if
		if val2 > 0 {
			fmt.Println("Больше нуля")
		} else {
			fmt.Println("Меньше нуля")
		}
		if val2 > 0 {
			fmt.Println("Больше нуля")
		}

		// ТЕМА: массив
		// [1,2,3,4,10]
		var mySlice []int = []int{1, 2, 3}
		fmt.Println("mySlice =", mySlice)
		slice := []int{1, 2, 3, 4, 5}
		//var slice []int
		fmt.Println("slice =", slice)

		fmt.Println("slice[0] =", slice[0])
		fmt.Println("slice len =", len(slice))

		// ТЕМА: цикл
		for i := 0; i < len(slice); i = i + 1 {
			fmt.Print("i =", i)
			fmt.Println(" slice[i] =", slice[i])
		}

		for i, v := range slice {
			fmt.Print("i =", i)
			fmt.Println(" v =", v)
		}

		var bigSlice []int

		for i := 2; i < 100; i = i + 2 {
			bigSlice = append(bigSlice, i)
		}
		fmt.Println(bigSlice)

		// ТЕМА: входные параметры
		fmt.Println("argument 0 =", os.Args[0])
		fmt.Println("argument 1 =", os.Args[1])
		firstArg, _ := strconv.Atoi(os.Args[1])
		secondArg, _ := strconv.Atoi(os.Args[2])
		fmt.Println("plus result =", firstArg+secondArg)
		fmt.Println("minus result =", firstArg-secondArg)

		fmt.Println("args len =", len(os.Args))

		var result string
		for index := 1; index < len(os.Args); index = index + 1 {

			result = result + os.Args[index]
		}
		fmt.Println("string result =", result)

		var intResult int
		for index := 1; index < len(os.Args); index = index + 1 {
			intArg, _ := strconv.Atoi(os.Args[index])
			intResult = intResult + intArg
		}
		fmt.Println("int result =", intResult)*/

	//fmt.Println(os.Args[1])

	// ТЕМА: функции
	/*b, err := fmt.Println("string")
	if err != nil {
		panic(err)
	}
	fmt.Println("b =", b)

	sumResult := sum(1, -9)
	fmt.Println(sumResult)

	var gipotenuza float64
	var err error
	gipotenuza, err = calculateGipotenuza(3, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ура, мы получили гипотенузу =", gipotenuza)*/

	// ТЕМА: простенькие задачи
	// площадь круга Pi*R*R
	res, err := squareCircle(5)
	if err != nil {
		panic(err)
	}
	fmt.Println("площадь круга =", res)

	err = calcSquare(7)
	if err != nil {
		panic(err)
	}

	// фиббоначи
	// функция должна вернуть n-ное число в последовательности фиббоначи
	var n uint64 = 20
	fib, err := getNFibbonachi(n)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%d число фиббоначи = %d", n, fib))

	// ТЕМА: решение первой задачи из leetcode https://leetcode.com/problems/two-sum/

	fmt.Println(reversArray([]int{0, 9, 8, 5}))
	fmt.Println(reversArray2([]int{1, 2, 3, 4, 5}))

	// ТЕМА: пакеты и области видимости
	// ТЕМА: ссылки

	mapExample()

	map1 := sliceToMap([]int{8, 88, 888, 8888})
	fmt.Println("Мапа из слайса", map1)
}

// 1 1 2 3 5 8
func getNFibbonachi(n uint64) (uint64, error) {
	if n == 0 {
		return 0, errors.New("n = 0")
	}
	var i uint64
	a := 0
	b := 1
	for i = 1; i < n; i = i + 1 {
		c := a
		a = b
		b = c + a
	}
	return uint64(b), nil
}

func sum(a int, b int) int {
	result := a + b
	return result
}

func sum2(a int, b int) (res int) {
	res = a + b
	return
}

// Для промяугольного треугольника
func calculateGipotenuza(catet1 int, catet2 int) (gipotenuza float64, err error) {
	if catet1 <= 0 {
		err = errors.New("invalid catet1")
		return
	}
	if catet2 <= 0 {
		err = errors.New("invalid catet2")
		return
	}
	gipotenuza2 := math.Pow(float64(catet1), 2) + math.Pow(float64(catet2), 2)
	gipotenuza = math.Sqrt(gipotenuza2)
	return
}

func calcSquare(v float64) (err error) {
	// calc circle
	res, err := squareCircle(v)
	if err != nil {
		return
	}
	fmt.Println("площадь круга =", res)

	// calc square
	res, err = squareSquare(v)
	if err != nil {
		return
	}
	fmt.Println("площадь квадрата =", res)
	return
}

func squareCircle(radius float64) (float64, error) {
	if radius <= 0 {
		err := errors.New("радиус меньше или равен нулю")
		return 0, err
	}
	square := math.Pi * math.Pow(radius, 2)
	return square, nil
}

func squareSquare(side float64) (square float64, err error) {
	if side <= 0 {
		err = errors.New("сторона меньше или равно нулю")
		return
	}
	square = math.Pow(side, 2)
	return
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i = i + 1 {
		for j := i + 1; j < len(nums); j = j + 1 {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func reversArray(slice []int) (output []int) {
	for i := len(slice) - 1; i >= 0; i = i - 1 {
		output = append(output, slice[i])
	}
	return output
}

func reversArray2(slice []int) []int {
	for i := 0; i < len(slice)/2; i = i + 1 {
		c := slice[i]
		slice[i] = slice[len(slice)-(i+1)]
		slice[len(slice)-(i+1)] = c
	}
	return slice
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

	ok, value := firstMap["Kazan"]
	if !ok {
		fmt.Println("Казани нет в списке")
	} else {
		fmt.Println("Казань есть в списке и равно", value)
	}
	firstMap["Kazan"] = true
	ok, value = firstMap["Kazan"]
	if !ok {
		fmt.Println("Казани нет в списке до сих пор")
	} else {
		fmt.Println("Казань есть теперь в списке и равно", value)
	}
	fmt.Println("Количество определенных городов в хэш мапе равно", len(firstMap))
}

func sliceToMap(slice []int) map[int]int {
	if len(slice) == 0 {
		return nil
	}
	map1 := make(map[int]int)
	for key, value := range slice {
		map1[value] = key
	}
	return map1
}
