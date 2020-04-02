package main

import (
	"testing"
	"time"
)

func TestGetSeconds(t *testing.T) {
	year := 2021
	month := time.January
	day := 1

	location, _ := time.LoadLocation("Europe/Stockholm")
	var SecondsTests = []struct {
		in   time.Time
		out  int
		name string
	}{
		{time.Date(year, month, day, 6, 0, 0, 0, location), 21600, "Normal case 6 oclock"},
		{time.Date(year, month, day, 12, 0, 0, 0, location), 43200, "Normal case 12 oclock"},
		{time.Date(year, month, day, 0, 0, 0, 0, location), 0, "Corner case midnight"},
		{time.Date(year, month, day, 0, 0, 1, 0, location), 1, "Corner case first second of day"},
		{time.Date(year, month, day, 23, 59, 59, 0, location), 86399, "Corner case 23:59:59"},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.name, func(t *testing.T) {
			s := toDaySeconds(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestFreeDates(t *testing.T) {
	location, _ := time.LoadLocation("Europe/Stockholm")
	var SecondsTests = []struct {
		in   time.Time
		out  bool
		name string
	}{
		{time.Date(2020, time.May, 16, 0, 0, 0, 0, location), true, "16th of May should be free"},
		{time.Date(2020, time.September, 19, 0, 0, 0, 0, location), true, "19th of September should be free"},
		{time.Date(2020, time.October, 30, 0, 0, 0, 0, location), true, "30th of October should be free"},
		{time.Date(2020, time.December, 31, 0, 0, 0, 0, location), true, "31th of December should be free"},
		{time.Date(2020, time.January, 1, 0, 0, 0, 0, location), true, "1st of January (New Years) should be free"},
		{time.Date(2020, time.December, 25, 0, 0, 0, 0, location), true, "25th of December (Christmas! :D) should be free"},
		{time.Date(2020, time.December, 26, 0, 0, 0, 0, location), true, "26th of December (Boxing day :( ) should be free"},
		{time.Date(2020, time.April, 2, 0, 0, 0, 0, location), false, "2th of April should be a working day"},
		{time.Date(2020, time.April, 4, 0, 0, 0, 0, location), true, "4th of April is a weekend"},
		{time.Date(2020, time.April, 5, 0, 0, 0, 0, location), true, "5th of April is a weekend"},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.name, func(t *testing.T) {
			s := isFreeDate(tt.in)
			if s != tt.out {
				t.Errorf("Got [%v] but expected [%v]", s, tt.out)
			}
		})
	}
}

func TestInRange(t *testing.T) {
	var SecondsTests = []struct {
		in          int
		min         int
		max         int
		out         bool
		description string
	}{
		{6, 5, 7, true, "5 <= 6 < 7, so true"},
		{6, 6, 7, true, "6 <= 6 < 7, so true"},
		{6, 5, 6, false, "5 <= 6 = 6, so false"},
		{5, 5, 5, false, "5 <= 5 = 5, so false.  This one is an edgy case, but since it does not occur in hour time ranges, it is fine. Another possible output would be an error."},
		{in: 5, min: 5, max: 3, out: false, description: "5 <= 5 < 3, so false. This could be another point of discussion, but we define that as accpted behavior since it does not happen for us."},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.description, func(t *testing.T) {
			s := closedOpenRange(tt.min, tt.max, tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestFindingInRangeIgnoringDay(t *testing.T) {
	location, _ := time.LoadLocation("Europe/Stockholm")
	year := 2021
	month := time.January
	day := 1

	var SecondsTests = []struct {
		in          time.Time
		out         Fee
		description string
	}{
		{time.Date(year, month, day, 16, 0, 0, 0, location), Increased, "16:00 is increased time"},
		{time.Date(year, month, day, 0, 0, 0, 0, location), Free, "00:00 is free time"},
		{time.Date(year, month, day, 5, 0, 0, 0, location), Free, "05:00 is free time"},
		{time.Date(year, month, day, 6, 0, 0, 0, location), Regular, "06:00 is regular time"},
		{time.Date(year, month, day, 7, 45, 0, 0, location), High, "07:45 is high time"},
		{time.Date(year, month, day, 8, 30, 0, 0, location), Increased, "08:30 is increased time"},
		{time.Date(year, month, day, 12, 0, 0, 0, location), Regular, "12:00 is regular time"},
		{time.Date(year, month, day, 17, 0, 0, 0, location), High, "17:00 is high time"},
		{time.Date(year, month, day, 18, 0, 0, 0, location), Increased, "18:00 is increased time"},
		{time.Date(year, month, day, 20, 0, 0, 0, location), Regular, "20:00 is regular time"},
		{time.Date(year, month, day, 23, 0, 0, 0, location), Free, "23:00 is free time"},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.description, func(t *testing.T) {
			s := findInRange(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestFindingFeeByTime(t *testing.T) {
	location, _ := time.LoadLocation("Europe/Stockholm")

	var SecondsTests = []struct {
		time        time.Time
		vehicle     VehicleType
		out         Fee
		description string
	}{
		{time.Date(2020, time.May, 16, 16, 0, 0, 0, location), Car, Free, "16th of April is a free day"},
		{time.Date(2020, time.May, 16, 16, 0, 0, 0, location), Motorbike, Free, "16th of April is a free day"},
		{time.Date(2020, time.April, 2, 16, 0, 0, 0, location), Motorbike, Free, "Motorbikes are always free"},
		{time.Date(2020, time.April, 2, 10, 0, 0, 0, location), Truck, Regular, "Normal pay (Regular) for a truck in the 2th of April"},
		{time.Date(2020, time.January, 12, 10, 0, 0, 0, location), Car, Free, "New years is a free day for all types of vehicles."},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Car, Regular, "April 2nd @ 10AM is Regular for Car"},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Motorbike, Free, "April 2nd @ 10AM is Free for Motorbike"},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Truck, Regular, "April 2nd @ 10AM is Regular for Truck"},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Tractor, Free, "April 2nd @ 10AM is Free for Tractor"},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Emergency, Free, "April 2nd @ 10AM is Free for Emergency"},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Diplomat, Free, "April 2nd @ 10AM is Free for Diplomat"},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Foreign, Free, "April 2nd @ 10AM is Free for Foreign"},
		{time.Date(2020, time.April, 7, 10, 0, 0, 0, location), Military, Free, "April 2nd @ 10AM is Free for Military"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Car, Free, "April 2nd @ 11PM is Free for Car"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Motorbike, Free, "April 2nd @ 11PM is Free for Motorbike"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Truck, Free, "April 2nd @ 11PM is Free for Truck"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Tractor, Free, "April 2nd @ 11PM is Free for Tractor"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Emergency, Free, "April 2nd @ 11PM is Free for Emergency"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Diplomat, Free, "April 2nd @ 11PM is Free for Diplomat"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Foreign, Free, "April 2nd @ 11PM is Free for Foreign"},
		{time.Date(2020, time.April, 7, 23, 0, 0, 0, location), Military, Free, "April 2nd @ 11PM is Free for Military"},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.description, func(t *testing.T) {
			s := getFee(tt.time, tt.vehicle).Fee
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
