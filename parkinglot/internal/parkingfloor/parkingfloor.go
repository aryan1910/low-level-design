package parkingfloor

import (
	"github.com/aryan1910/lowleveldesign/parkinglot/internal/parkingslot"
)

type ParkingFloor struct {
    FloorNumber     int
    ParkingLotID    string
    ParkingSlots    []*parkingslot.ParkingSlot
}