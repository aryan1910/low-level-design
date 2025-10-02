package ticket


import (
    "time"
    "github.com/aryan1910/lowleveldesign/parkinglot/internal/vehicle"
    "github.com/aryan1910/lowleveldesign/parkinglot/internal/parkingslot"
)


type Ticket struct {
    ID           string
    ParkingSlot  *parkingslot.ParkingSlot
    Vehicle      *vehicle.Vehicle
    EntryTime    time.Time
    ExitTime     *time.Time // nil if not exited
}