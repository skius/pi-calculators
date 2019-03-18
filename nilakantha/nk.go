package nilakantha

import "fmt"

var prefix = "Nilakantha:"

func CalculatePi() {
	var pi float64 = 3
	var divisorStart int64 = 2
	isPlus := true
	var steps uint64 = 0
	for {
		if isPlus {
			pi += 4/float64(divisorStart*(divisorStart+1)*(divisorStart+2))
		} else {
			pi -= 4/float64(divisorStart*(divisorStart+1)*(divisorStart+2))
		}
		divisorStart += 2
		isPlus = !isPlus
		steps++
		if steps % 100000000 == 0 {
			fmt.Println(prefix, pi)
		}
	}
}
