package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const (
		distance      = 62100000
		priceConst    = 20
		secondsPerDay = 86400
	)
	var companies = [3]string{"Virgin Galactic", "SpaceX", "Space Adventures"}
	var tripTypeName = [2]string{"One-way", "Round-trip"}

	fmt.Printf(
		"%-16v %4v %2v %4v\n=====================================\n",
		"Spaceline", "Days", "Trip type", "Price",
	)

	for count := 0; count < 10; count++ {
		station := companies[rand.Intn(2)]
		speed := rand.Intn(15) + 16
		tripType := rand.Intn(2) + 1
		price := (priceConst + speed) * tripType
		time := distance / speed / secondsPerDay // days

		fmt.Printf("%-16v %3v %-10v $%3v\n", station, time, tripTypeName[tripType-1], price)
	}
}
