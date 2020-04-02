package main

import (
	"testing"
	"time"
)

func TestCarWithVariousOverlaps(t *testing.T) {
	year := 2021
	month := time.April
	day := 1

	location, _ := time.LoadLocation("Europe/Stockholm")
	times := []time.Time{
		time.Date(year, month, day, 5, 0, 0, 0, location),
		time.Date(year, month, day, 6, 29, 0, 0, location),
		time.Date(year, month, day, 6, 59, 0, 0, location),
		time.Date(year, month, day, 7, 00, 0, 0, location),
		time.Date(year, month, day, 7, 01, 0, 0, location),
		time.Date(year, month, day, 8, 40, 0, 0, location),
		time.Date(year, month, day, 10, 00, 0, 0, location),
		time.Date(year, month, day, 10, 30, 0, 0, location),
		time.Date(year, month, day, 11, 20, 0, 0, location),
		time.Date(year, month, day, 11, 21, 0, 0, location),
		time.Date(year, month, day, 11, 22, 0, 0, location),
		time.Date(year, month, day, 12, 00, 0, 0, location),
		time.Date(year, month, day, 12, 10, 0, 0, location),
		time.Date(year, month, day, 12, 15, 0, 0, location),
		time.Date(year, month, day, 12, 30, 0, 0, location),
		time.Date(year, month, day, 14, 30, 0, 0, location),
	}

	expected := CollectFees(Regular, Increased, Regular, Regular, Regular, Regular)

	t.Run("Drove a lot, but only charged 6 times", func(t *testing.T) {
		s := CalculateFeesInADay(times, Car)
		if s != expected {
			t.Errorf("got %v, want %v", s, expected)
		}
	})
}

func TestCarHappyEasyDrive(t *testing.T) {
	year := 2021
	month := time.April
	day := 1

	location, _ := time.LoadLocation("Europe/Stockholm")
	times := []time.Time{
		time.Date(year, month, day, 7, 0, 0, 0, location),
		time.Date(year, month, day, 11, 00, 0, 0, location),
		time.Date(year, month, day, 11, 50, 0, 0, location),
		time.Date(year, month, day, 18, 30, 0, 0, location),
	}

	expected := CollectFees(High, Regular, Regular)

	t.Run("Drove a lot, but only charged 6 times", func(t *testing.T) {
		s := CalculateFeesInADay(times, Car)
		if s != expected {
			t.Errorf("got %v, want %v", s, expected)
		}
	})
}

func TestCarDrivingEachHourThenBlowingThreshold(t *testing.T) {
	year := 2021
	month := time.April
	day := 1

	location, _ := time.LoadLocation("Europe/Stockholm")
	times := []time.Time{
		time.Date(year, month, day, 6, 0, 0, 0, location),
		time.Date(year, month, day, 7, 0, 0, 0, location),
		time.Date(year, month, day, 8, 30, 0, 0, location),
		time.Date(year, month, day, 9, 30, 0, 0, location),
		time.Date(year, month, day, 10, 40, 0, 0, location),
		time.Date(year, month, day, 11, 40, 0, 0, location),
	}

	expected := 60 //expects the car to blow the maximum count in a day

	t.Run("Drove a lot, but only charged 6 times", func(t *testing.T) {
		s := CalculateFeesInADay(times, Car)
		if s != expected {
			t.Errorf("got %v, want %v", s, expected)
		}
	})
}

func TestAllAllegedlyFreeVehiclesShouldPayNoToll(t *testing.T) {
	year := 2021
	month := time.April
	day := 1

	location, _ := time.LoadLocation("Europe/Stockholm")
	times := []time.Time{
		time.Date(year, month, day, 6, 0, 0, 0, location),
		time.Date(year, month, day, 7, 0, 0, 0, location),
		time.Date(year, month, day, 8, 30, 0, 0, location),
		time.Date(year, month, day, 9, 30, 0, 0, location),
		time.Date(year, month, day, 10, 40, 0, 0, location),
		time.Date(year, month, day, 11, 40, 0, 0, location),
	}

	expected := 0 //this type of vehicle should always be free, no matter how many drives

	expectedFreeVehicles := []VehicleType{Motorbike, Tractor, Emergency, Diplomat, Foreign, Military}

	t.Run("Drove a lot, but only charged 6 times", func(t *testing.T) {
		for _, freeVehicle := range expectedFreeVehicles {
			s := CalculateFeesInADay(times, freeVehicle)
			if s != expected {
				t.Errorf("got %v, want %v", s, expected)
			}

		}
	})
}

func TestCarHappyEasyDriveManyDaysOfTheWeek(t *testing.T) {
	year := 2020
	month := time.April
	day := 1

	location, _ := time.LoadLocation("Europe/Stockholm")
	times := []time.Time{
		time.Date(year, month, day, 7, 0, 0, 0, location),
		time.Date(year, month, day, 11, 00, 0, 0, location),
		time.Date(year, month, day, 11, 50, 0, 0, location),
		time.Date(year, month, day, 18, 30, 0, 0, location),

		time.Date(year, month, day+1, 7, 0, 0, 0, location),
		time.Date(year, month, day+1, 11, 00, 0, 0, location),
		time.Date(year, month, day+1, 11, 50, 0, 0, location),
		time.Date(year, month, day+1, 18, 30, 0, 0, location),

		time.Date(year, month, day+2, 7, 0, 0, 0, location),
		time.Date(year, month, day+2, 11, 00, 0, 0, location),
		time.Date(year, month, day+2, 11, 50, 0, 0, location),
		time.Date(year, month, day+2, 18, 30, 0, 0, location),

		time.Date(year, month, 4, 7, 0, 0, 0, location),
		time.Date(year, month, 4, 11, 00, 0, 0, location),
		time.Date(year, month, 4, 11, 50, 0, 0, location),
		time.Date(year, month, 4, 18, 30, 0, 0, location),

		time.Date(year, month, 5, 7, 0, 0, 0, location),
		time.Date(year, month, 5, 11, 00, 0, 0, location),
		time.Date(year, month, 5, 11, 50, 0, 0, location),
		time.Date(year, month, 5, 18, 30, 0, 0, location),
	}

	//only 3 times as we have 2 weekend days
	expected := 3 * CollectFees(High, Regular, Regular)

	t.Run("Drove a lot, but only charged 6 times", func(t *testing.T) {
		s := CalculateAllFees(times, Car)
		if s != expected {
			t.Errorf("Got [%v], but wanted [%v]", s, expected)
		}
	})

}

func TestCarHappyEasyDriveManyDaysOfTheWeekAndWeekend(t *testing.T) {
	year := 2020
	month := time.April
	day := 1

	location, _ := time.LoadLocation("Europe/Stockholm")
	times := []time.Time{
		time.Date(year, month, day, 7, 0, 0, 0, location),
		time.Date(year, month, day, 11, 00, 0, 0, location),
		time.Date(year, month, day, 11, 50, 0, 0, location),
		time.Date(year, month, day, 18, 30, 0, 0, location),

		time.Date(year, month, day+1, 7, 0, 0, 0, location),
		time.Date(year, month, day+1, 11, 00, 0, 0, location),
		time.Date(year, month, day+1, 11, 50, 0, 0, location),
		time.Date(year, month, day+1, 18, 30, 0, 0, location),

		time.Date(year, month, day+2, 7, 0, 0, 0, location),
		time.Date(year, month, day+2, 11, 00, 0, 0, location),
		time.Date(year, month, day+2, 11, 50, 0, 0, location),
		time.Date(year, month, day+2, 18, 30, 0, 0, location),
	}

	expected := 3 * CollectFees(High, Regular, Regular)

	t.Run("Drove a lot, but only charged 6 times", func(t *testing.T) {
		s := CalculateAllFees(times, Car)
		if s != expected {
			t.Errorf("Got [%v], but wanted [%v]", s, expected)
		}
	})

}

func TestCarDrivingOnlyOnHolidays(t *testing.T) {
	year := 2021

	location, _ := time.LoadLocation("Europe/Stockholm")
	times := []time.Time{
		time.Date(year, time.December, 25, 7, 0, 0, 0, location),   //xmas!
		time.Date(year, time.December, 25, 11, 0, 0, 0, location),  //xmas!
		time.Date(year, time.December, 25, 11, 50, 0, 0, location), //xmas!
		time.Date(year, time.December, 25, 18, 0, 0, 0, location),  //xmas!

		time.Date(year, time.January, 1, 7, 0, 0, 0, location),   //new years
		time.Date(year, time.January, 1, 11, 00, 0, 0, location), //new years
		time.Date(year, time.January, 1, 11, 50, 0, 0, location), //new years
		time.Date(year, time.January, 1, 18, 30, 0, 0, location), //new years

		time.Date(year, time.May, 16, 7, 0, 0, 0, location),   //ordered holiday
		time.Date(year, time.May, 16, 11, 00, 0, 0, location), //ordered holiday
		time.Date(year, time.May, 16, 11, 50, 0, 0, location), //ordered holiday
		time.Date(year, time.May, 16, 18, 30, 0, 0, location), //ordered holiday

		time.Date(year, time.September, 19, 7, 0, 0, 0, location),   //ordered holiday
		time.Date(year, time.September, 19, 11, 00, 0, 0, location), //ordered holiday
		time.Date(year, time.September, 19, 11, 50, 0, 0, location), //ordered holiday
		time.Date(year, time.September, 19, 18, 30, 0, 0, location), //ordered holiday

		time.Date(year, time.October, 30, 7, 0, 0, 0, location),   //ordered holiday
		time.Date(year, time.October, 30, 11, 00, 0, 0, location), //ordered holiday
		time.Date(year, time.October, 30, 11, 50, 0, 0, location), //ordered holiday
		time.Date(year, time.October, 30, 18, 30, 0, 0, location), //ordered holiday

		time.Date(year, time.December, 31, 7, 0, 0, 0, location),   //ordered holiday
		time.Date(year, time.December, 31, 11, 00, 0, 0, location), //ordered holiday
		time.Date(year, time.December, 31, 11, 50, 0, 0, location), //ordered holiday
		time.Date(year, time.December, 31, 18, 30, 0, 0, location), //ordered holiday
	}

	expected := 0 //only holidays!

	t.Run("Drove a lot, but only charged 6 times", func(t *testing.T) {
		s := CalculateAllFees(times, Car)
		if s != expected {
			t.Errorf("Got [%v], but wanted [%v]", s, expected)
		}
	})

}
