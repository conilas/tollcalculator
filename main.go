package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	car := flag.Int("vehicle", 0, "The vehicle type. The range goes from 0 to 7")
	timestamps := flag.String("timestamps", "", "Unix timestamps, coma separated. E.g: ")
	flag.Parse()

	parsedCar, ok := ParseType(*car)
	if !ok {
		log.Fatalf("Could not parse argument [%v]. Invalid type in range 0-7.", *car)
		return
	}

	splittedTimes := strings.Split(*timestamps, ",")
	parsedTimes, err := MapToTimestamps(splittedTimes)
	if err != nil {
		log.Fatalf("Could not parse argument [%v]. Invalid coma separated unix times.", *timestamps)
		return
	}

	log.Printf("Your total collected tolls are [%v]. The times (human-readable) for tolls are [%v]. Please note that we use GMT+2/Swedish time to calculate tolls.",
		CalculateAllFees(parsedTimes, parsedCar), parsedTimes)
}
