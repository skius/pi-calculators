package main

import (
	"github.com/skius/pi-calculators/gregoryleibniz"
	"github.com/skius/pi-calculators/nilakantha"
)

func main() {
	go gregoryleibniz.CalculatePi()
	nilakantha.CalculatePi()
}
