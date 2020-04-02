package main

import (
	"log"
	//"sort"
	"time"
)

func toFee(v VehicleType) func(time.Time) ChargeableFee {
	return func(t time.Time) ChargeableFee {
		return getFee(t, v)
	}
}

type ByTime []ChargeableFee

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }

func getSingleFeeFromList(charges []ChargeableFee) Fee {
	return FindFirstThat(OrderChargeableFees(charges), func(fee ChargeableFee) bool {
		return fee.Fee != Free
	}, Free)
}

func CalculateFeesInADay(times []time.Time, v VehicleType) int {
	chargeableFees := OrderChargeableFees(MapToFee(times, toFee(v)))
	mappedFees := ToMapFromHoursToFees(chargeableFees)
	uniqueFeesPerHour := ReduceFromHourFeeToSingleFeeList(mappedFees, getSingleFeeFromList)
	collectedFeesOnDay := ReduceFees(uniqueFeesPerHour, IncrementFee)

	log.Printf("The accumulated value for the day is [%v]. The hourly map is [%v]. ", collectedFeesOnDay, mappedFees)

	return Min(collectedFeesOnDay, 60)
}

func CalculateFeesInADayForCar(v VehicleType) func([]time.Time) int {
	return func(times []time.Time) int {
		return CalculateFeesInADay(times, v)
	}
}

func CalculateAllFees(times []time.Time, v VehicleType) int {
	mapWithDates := MapFromTimesToDailyValues(times)
	reducedValues := ReduceTimesToDailyFees(mapWithDates, CalculateFeesInADayForCar(v))
	totalFees := ReduceDailyMapToFees(reducedValues, Sum)

	log.Printf("The values for each day are [%v] for car of type [%v]. The sum of all fees is [%v]", reducedValues, v, totalFees)

	return totalFees
}
