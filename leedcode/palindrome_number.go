package leedcode

import (
	"fmt"
)

func palindromeNumberStr(x int) bool {
	if x < 0 {
		return false
	}
	str := fmt.Sprintf("%d", x)
	for i := 0; i < len(str)/2; i = i + 1 {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func palindromeNumberInt(x int) bool {
	if x < 0 {
		return false
	}
	z := 0
	for y := x; y > 0; {
		z = z*10 + y%10
		y = y / 10
	}
	if z == x {
		return true
	}
	return false
}
