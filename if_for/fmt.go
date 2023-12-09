package if_for

import (
	"fmt"
)

func ExamplePrintln() {
	// Он возвращает количество записанных байтов и все обнаруженные ошибки записи.
	b, err := fmt.Println("string")
	if err != nil {
		panic(err)
	}
	fmt.Println("количество записанных байтов в Println =", b)
}
