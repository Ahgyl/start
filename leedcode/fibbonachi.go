package leedcode

import (
	"fmt"
)

func ExampleFibbonachi() {
	// функция должна вернуть n-ное число в последовательности фиббоначи
	var n uint64 = 20
	fib := getNFibbonachi(n)
	fmt.Println(fmt.Sprintf("%d число фиббоначи = %d", n, fib))
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597...
func getNFibbonachi(n uint64) uint64 {
	if n == 0 {
		return 0
	}
	var i uint64
	a := 0
	b := 1
	for i = 1; i < n; i = i + 1 {
		c := a
		a = b
		b = c + a
	}
	return uint64(b)
}
