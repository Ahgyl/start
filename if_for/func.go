package if_for

import (
	"fmt"
)

func ExampleFunc() {
	sumResult := sum(1, -9)
	fmt.Println("Результат сложения двух числе 1 и -9 функцией sum", sumResult)
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
