package leedcode

var (
	romanMap = map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
)

// IVI
func RomanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			if i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
				// s[i] == I
				// s[i+1] == V or X
				key := string(s[i]) + string(s[i+1])
				res = res + romanMap[key]
				i++
			} else {
				res = res + romanMap[string(s[i])]
			}
		} else if s[i] == 'X' {
			if i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
				// s[i] == X
				// s[i+1] == L or C
				key := string(s[i]) + string(s[i+1])
				res = res + romanMap[key]
				i++
			} else {
				res = res + romanMap[string(s[i])]
			}
		} else if s[i] == 'C' {
			if i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
				// s[i] == C
				// s[i+1] == D or M
				key := string(s[i]) + string(s[i+1])
				res = res + romanMap[key]
				i++
			} else {
				res = res + romanMap[string(s[i])]
			}
		} else {
			res = res + romanMap[string(s[i])]
		}

	}
	return res
}
