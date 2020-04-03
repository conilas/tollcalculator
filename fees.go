package main

import (
	"github.com/rickar/cal"
	"time"
)

type Fee int

const (
	Free      Fee = 0
	Regular   Fee = 8
	Increased Fee = 13
	High      Fee = 18
)

type FeeRange struct {
	StartingAt int
	EndingAt   int
	Value      Fee
}

type ChargeableFee struct {
	Fee  Fee
	Hour int
	Time time.Time
}

type InRangePredicate func(ran FeeRange) bool

var FeeRanges = []FeeRange{
	{StartingAt: toSecondsFromHours(6), EndingAt: toSeconds(6, 30), Value: Regular},
	{StartingAt: toSeconds(6, 30), EndingAt: toSecondsFromHours(7), Value: Increased},
	{StartingAt: toSecondsFromHours(7), EndingAt: toSeconds(8, 30), Value: High},
	{StartingAt: toSeconds(8, 30), EndingAt: toSecondsFromHours(9), Value: Increased},
	{StartingAt: toSecondsFromHours(9), EndingAt: toSecondsFromHours(16), Value: Regular},
	{StartingAt: toSecondsFromHours(16), EndingAt: toSeconds(16, 30), Value: Increased},
	{StartingAt: toSeconds(16, 30), EndingAt: toSecondsFromHours(18), Value: High},
	{StartingAt: toSecondsFromHours(18), EndingAt: toSeconds(18, 30), Value: Increased},
	{StartingAt: toSeconds(18, 30), EndingAt: toSecondsFromHours(21), Value: Regular},
}

func toSecondsFromHours(hours int) int {
	return toSeconds(hours, 0)
}

func toSeconds(hours int, minutes int) int {
	return (hours * 60 * 60) + (minutes * 60)
}

func toDaySeconds(t time.Time) int {
	year, month, day := t.Date()
	currentStartDate := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return int(t.Sub(currentStartDate).Seconds())
}

func isFreeDate(t time.Time) bool {
	c := cal.NewCalendar()

	cal.AddSwedishHolidays(c)

	//these are all special cases, but let's consider them as holidays as well
	c.AddHoliday(
		cal.NewHoliday(time.May, 16),
		cal.NewHoliday(time.September, 19),
		cal.NewHoliday(time.October, 30),
		cal.NewHoliday(time.December, 31),
	)

	return !c.IsWorkday(t)
}

func closedOpenRange(min int, max int, value int) bool {
	return min <= value && value < max
}

func testRangeWith(seconds int) InRangePredicate {
	return func(ran FeeRange) bool {
		return closedOpenRange(ran.StartingAt, ran.EndingAt, seconds)
	}
}

func findInRange(t time.Time) Fee {
	seconds := toDaySeconds(t)
	return FilterRange(FeeRanges, testRangeWith(seconds))
}

func IncrementFee(value int, fee Fee) int {
	return int(fee) + value
}

func CollectFees(fees ...Fee) int {
	total := 0
	for _, fee := range fees {
		total += int(fee)
	}
	return total
}

func getFee(t time.Time, v VehicleType) ChargeableFee {
	if isFreeDate(t) || IsFreeVehicle(v) {
		return ChargeableFee{Fee: Free, Hour: t.Hour(), Time: t}
	}

	return ChargeableFee{Fee: findInRange(t), Hour: t.Hour(), Time: t}
}
