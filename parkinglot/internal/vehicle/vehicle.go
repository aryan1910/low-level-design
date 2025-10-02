package vehicle

import "fmt"

type VehicleType string

const (
	VehicleTypeCar   VehicleType = "CAR"
	VehicleTypeBike  VehicleType = "BIKE"
	VehicleTypeTruck VehicleType = "TRUCK"
)

type Vehicle struct {
	Type  VehicleType
	RegNo string
	Color string
}

func NewVehicle(vt VehicleType, regNo, color string) *Vehicle {
    return &Vehicle{
        Type:  vt,
        RegNo: regNo,
        Color: color,
    }
}

func ParseVehicleType(s string) (VehicleType, error) {
	switch s {
	case "CAR":
		return VehicleTypeCar, nil
	case "BIKE":
		return VehicleTypeBike, nil
	case "TRUCK":
		return VehicleTypeTruck, nil
	default:
		return "", fmt.Errorf("invalid vehicle type: %s", s)
	}
}
