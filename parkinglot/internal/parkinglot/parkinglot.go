package parkinglot

import (
	"fmt"
	"strings"

	"github.com/aryan1910/lowleveldesign/parkinglot/internal/parkingfloor"
	"github.com/aryan1910/lowleveldesign/parkinglot/internal/parkingslot"
	"github.com/aryan1910/lowleveldesign/parkinglot/internal/vehicle"
)

type ParkingLot struct {
	ID                string
	NoOfFloors        int
	NoOfSlotsPerFloor int
	ParkingFloors     []*parkingfloor.ParkingFloor
}

func NewParkingLot(id string, noOfFloors, noOfSlotsPerFloor int) *ParkingLot {
	floorList := make([]*parkingfloor.ParkingFloor, 0, noOfFloors)
	for i := 0; i < noOfFloors; i++ {
		slotList := make([]*parkingslot.ParkingSlot, 0, noOfSlotsPerFloor)
		for j := 0; j < noOfSlotsPerFloor; j++ {
			vt := vehicle.VehicleTypeCar
			if j == 0 {
				vt = vehicle.VehicleTypeTruck
			} else if j == 1 || j == 2 {
				vt = vehicle.VehicleTypeBike
			}
			slot := &parkingslot.ParkingSlot{
				ID:              fmt.Sprintf("%d", j+1),
				VehicleType:     vt,
				Status:          parkingslot.StatusAvailable,
				ParkingLotID:    id,
				ParkingFloorNum: i + 1,
			}
			slotList = append(slotList, slot)
		}
		floor := &parkingfloor.ParkingFloor{
			FloorNumber:  i + 1,
			ParkingLotID: id,
			ParkingSlots: slotList,
		}
		floorList = append(floorList, floor)
	}
	return &ParkingLot{
		ID:                id,
		NoOfFloors:        noOfFloors,
		NoOfSlotsPerFloor: noOfSlotsPerFloor,
		ParkingFloors:     floorList,
	}
}

func (pl *ParkingLot) DisplayFreeCount(vehicleType vehicle.VehicleType) {
	for _, floor := range pl.ParkingFloors {
		count := 0
		for _, slot := range floor.ParkingSlots {
			if slot.VehicleType == vehicleType && slot.Status == parkingslot.StatusAvailable {
				count++
			}
		}
		fmt.Printf("No. of free slots for %s on Floor %d: %d\n", vehicleType, floor.FloorNumber, count)
	}
}

func (pl *ParkingLot) DisplayFreeSlots(vehicleType vehicle.VehicleType) {
	for _, floor := range pl.ParkingFloors {
		var slotIDs []string
		for _, slot := range floor.ParkingSlots {
			if slot.VehicleType == vehicleType && slot.Status == parkingslot.StatusAvailable {
				slotIDs = append(slotIDs, slot.ID)
			}
		}
		fmt.Printf("Free slots for %s on Floor %d: %s\n", vehicleType, floor.FloorNumber, strings.Join(slotIDs, ","))
	}
}

func (pl *ParkingLot) DisplayOccupiedSlots(vehicleType vehicle.VehicleType) {
	for _, floor := range pl.ParkingFloors {
		var slotIDs []string
		for _, slot := range floor.ParkingSlots {
			if slot.VehicleType == vehicleType && slot.Status == parkingslot.StatusOccupied {
				slotIDs = append(slotIDs, slot.ID)
			}
		}
		fmt.Printf("Occupied slots for %s on Floor %d: %s\n", vehicleType, floor.FloorNumber, strings.Join(slotIDs, ","))
	}
}

func (pl *ParkingLot) ParkVehicle(v *vehicle.Vehicle) (ticketID string, err error) {
    for _, floor := range pl.ParkingFloors {
        for _, slot := range floor.ParkingSlots {
            if slot.VehicleType == v.Type && slot.Status == parkingslot.StatusAvailable {
                slot.Status = parkingslot.StatusOccupied
                slot.Vehicle = v
                ticketID := fmt.Sprintf("PR%s_%d_%s", pl.ID, floor.FloorNumber, slot.ID)
                return ticketID, nil
            }
        }
    }
    return "", fmt.Errorf("no available slot for vehicle type: %s", v.Type)
}