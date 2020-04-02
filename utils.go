package main

import (
	"sort"
	"strconv"
	"time"
)

func FilterRange(ranges []FeeRange, f func(ran FeeRange) bool) Fee {
	for _, v := range ranges {
		if f(v) {
			return v.Value
		}
	}
	return Free
}

func MapToFee(times []time.Time, f func(t time.Time) ChargeableFee) []ChargeableFee {
	vsm := make([]ChargeableFee, len(times))
	for i, v := range times {
		vsm[i] = f(v)
	}
	return vsm
}

func ToMapFromHoursToFees(chargeableFees []ChargeableFee) map[time.Time][]ChargeableFee {
	return toMapFromHoursToFeesRec(chargeableFees, make(map[time.Time][]ChargeableFee))
}

func isEmpty(fees []ChargeableFee) bool {
	return len(fees) == 0
}

func head(chargeableFees []ChargeableFee) ChargeableFee {
	return chargeableFees[0]
}

func tail(chargeableFees []ChargeableFee) []ChargeableFee {
	return chargeableFees[1:]
}

func toMapFromHoursToFeesRec(fees []ChargeableFee, values map[time.Time][]ChargeableFee) map[time.Time][]ChargeableFee {
	if isEmpty(fees) {
		return values
	}

	head := head(fees)
	tail := tail(fees)

	for key, filledValues := range values {
		if key.Add(time.Hour).After(head.Time) {
			values[key] = append(filledValues, head)
			return toMapFromHoursToFeesRec(tail, values)
		}
	}

	values[head.Time] = []ChargeableFee{head}
	return toMapFromHoursToFeesRec(tail, values)
}

func MapFromTimesToDailyValues(times []time.Time) map[string][]time.Time {
	accMap := make(map[string][]time.Time)
	for _, timeValue := range times {
		year, month, day := timeValue.Date()
		strTime := strconv.Itoa(day) + "-" + month.String() + "-" + strconv.Itoa(year)
		if accMap[strTime] != nil {
			accMap[strTime] = append(accMap[strTime], timeValue)
		} else {
			accMap[strTime] = []time.Time{timeValue}
		}
	}
	return accMap
}

func ReduceTimesToDailyFees(values map[string][]time.Time, mapper func([]time.Time) int) map[string]int {
	finalMap := make(map[string]int)
	for key, values := range values {
		finalMap[key] = mapper(values)
	}
	return finalMap
}

func ReduceFromHourFeeToSingleFeeList(hourFees map[time.Time][]ChargeableFee, mapper func([]ChargeableFee) Fee) []Fee {
	vsm := make([]Fee, len(hourFees))
	for _, value := range hourFees {
		vsm = append(vsm, mapper(value))
	}
	return vsm
}

func FindFirstThat(fees []ChargeableFee, predicate func(ChargeableFee) bool, fallback Fee) Fee {
	for _, v := range fees {
		if predicate(v) {
			return v.Fee
		}
	}
	return fallback
}

func ReduceDailyMapToFees(values map[string]int, acc func(int, int) int) int {
	accumulated := 0
	for _, v := range values {
		accumulated = acc(accumulated, v)
	}
	return accumulated

}

func ReduceFees(fees []Fee, acc func(int, Fee) int) int {
	accumulated := 0
	for _, v := range fees {
		accumulated = acc(accumulated, v)
	}
	return accumulated
}

func OrderChargeableFees(charges []ChargeableFee) []ChargeableFee {
	sort.Sort(ByTime(charges))
	return charges
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Sum(x, y int) int {
	return x + y
}
