package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
	"time"
)

func MapToTimestamps(unixTimes []string) ([]time.Time, error) {
	parsedTimes := []time.Time{}
	for _, value := range unixTimes {
		parsed, err := parseUnixTime(value)
		parsedTimes = append(parsedTimes, parsed)
		if err != nil {
			return parsedTimes, err
		}
	}
	return parsedTimes, nil

}

func parseUnixTime(unixTimeStamp string) (time.Time, error) {
	timeParsed, err := strconv.ParseInt(unixTimeStamp, 10, 64)

	if err != nil {
		return time.Now(), err
	}

	return time.Unix(timeParsed, 0), nil
}

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

	log.Printf("Your total collected tolls are [%v]. The times (human-readable) for tools are [%v]. Please note that we use GMT+2/Swedish time to calculate tolls.",
		CalculateAllFees(parsedTimes, parsedCar), parsedTimes)
}
