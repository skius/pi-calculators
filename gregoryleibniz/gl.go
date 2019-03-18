package gregoryleibniz

import "fmt"

var prefix = "Gregory-Leibniz:"

func CalculatePi() {
	var fourthPi float64 = 1
	var divisor int64 = 3
	isPlus := false
	var steps uint64 = 0
	for {
		if isPlus {
			fourthPi += 1/float64(divisor)
		} else {
			fourthPi -= 1/float64(divisor)
		}
		divisor += 2
		isPlus = !isPlus
		steps++
		if steps % 100000000 == 0 {
			fmt.Println(prefix, 4*fourthPi)
		}
	}
}
