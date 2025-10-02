package parkingslot

import (
	"github.com/aryan1910/lowleveldesign/parkinglot/internal/vehicle"
)

type Status string

const (
	StatusAvailable Status = "AVAILABLE"
	StatusOccupied  Status = "OCCUPIED"
)

type ParkingSlot struct {
	ID              string
	VehicleType     vehicle.VehicleType
	Status          Status
	ParkingLotID    string
	ParkingFloorNum int
	Vehicle         *vehicle.Vehicle
}
