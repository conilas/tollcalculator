package main

const (
	Car VehicleType = iota
	Truck
	Motorbike
	Tractor
	Emergency
	Diplomat
	Foreign
	Military
)

type VehicleType int
type Vehicle struct {
	Type VehicleType
	Free bool
}

var Vehicles = []Vehicle{
	{Type: Motorbike, Free: true},
	{Type: Car, Free: false},
	{Type: Truck, Free: false},
	{Type: Tractor, Free: true},
	{Type: Emergency, Free: true},
	{Type: Diplomat, Free: true},
	{Type: Foreign, Free: true},
	{Type: Military, Free: true},
}

func IsFreeVehicle(v VehicleType) bool {
	for _, value := range Vehicles {
		if value.Type == v {
			return value.Free
		}
	}
	//maybe return error instead?
	return false
}

func ParseType(value int) (VehicleType, bool) {
	if value >= int(Car) && value <= int(Military) {
		return VehicleType(value), true
	}
	return Car, false
}
