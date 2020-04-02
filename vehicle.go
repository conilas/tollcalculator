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
	Vehicle{Type: Motorbike, Free: true},
	Vehicle{Type: Car, Free: false},
	Vehicle{Type: Truck, Free: false},
	Vehicle{Type: Tractor, Free: true},
	Vehicle{Type: Emergency, Free: true},
	Vehicle{Type: Diplomat, Free: true},
	Vehicle{Type: Foreign, Free: true},
	Vehicle{Type: Military, Free: true},
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
