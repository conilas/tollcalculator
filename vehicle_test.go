package main

import (
	"testing"
)

func TestCheckForFreeCars(t *testing.T) {
	var SecondsTests = []struct {
		in   VehicleType
		out  bool
		name string
	}{
		{Motorbike, true, "Motorbikes are free vehicles"},
		{Car, false, "Cars are NOT free cars"},
		{Truck, false, "Trucks are NOT free vehicles"},
		{Tractor, true, "Tractors are free vehicles"},
		{Emergency, true, "Emergencys are free vehicles"},
		{Diplomat, true, "Diplomats are free vehicles"},
		{Foreign, true, "Foreigns are free vehicles"},
		{Military, true, "Militarys are free vehicles"},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.name, func(t *testing.T) {
			s := IsFreeVehicle(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestParsingValues(t *testing.T) {
	var SecondsTests = []struct {
		in      int
		out     VehicleType
		correct bool
		name    string
	}{
		{0, Car, true, "Cars are OK"},
		{1, Truck, true, "Trucks are OK"},
		{2, Motorbike, true, "Motorbikes are OK"},
		{3, Tractor, true, "Tractors are OK"},
		{4, Emergency, true, "Emergencys are OK"},
		{5, Diplomat, true, "Diplomats are OK"},
		{6, Foreign, true, "Foreigns are OK"},
		{7, Military, true, "Militarys are OK"},
		{8, Car, false, "Invalid input"},
		{9, Car, false, "Invalid Input"},
		{10, Car, false, "Invalid Input"},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.name, func(t *testing.T) {
			car, correct := ParseType(tt.in)
			if correct != tt.correct {
				t.Errorf("got [%v], want [%v]", correct, tt.correct)
			}

			if car != tt.out {
				t.Errorf("got [%v], want [%v]", car, tt.out)
			}
		})
	}
}
