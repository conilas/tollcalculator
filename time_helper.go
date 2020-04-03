package main

import (
	"strconv"
	"time"
)

func MapToTimestamps(unixTimes []string) ([]time.Time, error) {
	parsedTimes := []time.Time{}
	for _, value := range unixTimes {
		parsed, err := parseUnixTime(value)
		parsedTimes = append(parsedTimes, parsed)
		if err != nil {
			return make([]time.Time, 0), err
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
