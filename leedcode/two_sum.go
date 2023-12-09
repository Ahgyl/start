package leedcode

import (
	"fmt"
)

func ExampleTwoSum() {
	// ТЕМА: решение первой задачи из leetcode https://leetcode.com/problems/two-sum/
	fmt.Println("twoSum =", twoSum2([]int{2, 7, 11, 15}, 9))
}

// Дан массив целых чисел nums и целое число target, функция возвращает индексы двух чисел так,
// чтобы их сумма была равна target .
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

// Работает быстрее за счет мапы
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for key, value := range nums {
		a := target - value
		b, ok := m[a]
		if ok {
			if key == b {
				continue
			}
			return []int{key, b}
		} else {
			m[value] = key
		}
	}
	return nil
}
