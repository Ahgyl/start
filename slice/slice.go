package slice

import (
	"fmt"
)

func Example() {
	fmt.Println("===Package slice===")
	// ТЕМА: массив
	var mySlice = []int{1, 2, 3}
	fmt.Println("слайс через var =", mySlice)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("слайс через :=", slice)

	fmt.Println("slice[0] =", slice[0])
	fmt.Println("slice len =", len(slice))

	var sliceAppend []int
	fmt.Println("Пустой sliceAppend", sliceAppend)
	sliceAppend = append(sliceAppend, 1, 2)
	fmt.Println("После добавления двух элементов", sliceAppend)
	for i := 10; i < 20; i = i + 1 {
		sliceAppend = append(sliceAppend, i)
	}
	fmt.Println("После append в цикле sliceAppend", sliceAppend)
}

// Запись слайса в обратном порядке
// 1. создавая новый слайс output
func ReversArray(slice []int) (output []int) {
	for i := len(slice) - 1; i >= 0; i = i - 1 {
		output = append(output, slice[i])
	}
	return output
}

// 2. оставляя тот же слайс
// Во входном массиве меняем местами первое значение с последним, второе с предпоследним и т.д.
func ReversArray2(slice []int) []int {
	for i := 0; i < len(slice)/2; i = i + 1 {
		c := slice[i]
		slice[i] = slice[len(slice)-(i+1)]
		slice[len(slice)-(i+1)] = c
	}
	return slice
}

// Перевод слайса в мапу (значение из слайса становится ключом в мапе, а индекс из слайса становится значением в мапе)
func SliceToMap(slice []int) map[int]int {
	if len(slice) == 0 {
		return nil
	}
	map1 := make(map[int]int)
	for key, value := range slice {
		map1[value] = key
	}
	return map1
}

// Сортировка пузырьком.
func BubbleSort(slice []int) []int {
	// Этот цикл запускает меньший цикл снова и снова, пока полностью все цифры не будут отсортированы по возрастанию
	for j := 0; j < len(slice); j = j + 1 {
		// Создали флаг, который сигнализирует о том что перестановка произошла.
		// В противном случае сортировка останавливается, по причине того что всё расставили по местам.
		var isSwap bool
		// Меньший цикл: проходится один раз по числам и выстраевает все на своем пути в одном направлении
		for i := 0; i < len(slice)-j-1; i = i + 1 {
			if slice[i] > slice[i+1] {
				c := slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = c
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
	return slice
}
