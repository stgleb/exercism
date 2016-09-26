package raindrops

import "fmt"

const testVersion = 2

var convert = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func check(value, number int) string {
	result := ""

	for {
		if value % number != 0 {
			break
		} else {
			result += convert[number]
			value = value / number
			break
		}
	}

	return result
}

func Convert(value int) string {
	result := check(value, 3)
	result += check(value, 5)
	result += check(value, 7)

	if len(result) == 0 || value < 3 {
		return fmt.Sprintf("%d", value)
	}

	return result
}