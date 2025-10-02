package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aryan1910/lowleveldesign/parkinglot/internal/parkinglot"
	"github.com/aryan1910/lowleveldesign/parkinglot/internal/parkingslot"
	"github.com/aryan1910/lowleveldesign/parkinglot/internal/vehicle"
)

type ParkingLotManager struct {
	parkingLots map[string]*parkinglot.ParkingLot
	primaryId   string
}

func main() {
	manager := ParkingLotManager{
		parkingLots: make(map[string]*parkinglot.ParkingLot),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(input)
		if len(parts) == 0 {
			fmt.Println("invalid command")~
			continue
		}

		cmd := parts[0]
		switch cmd {
		case "create_parking_lot":
			err := manager.create(parts)
			if err != nil {
				fmt.Println(err)
			}
		case "park_vehicle":
			err := manager.parkVehicle(parts)
			if err != nil {
				fmt.Println(err)
			}
		case "unpark_vehicle":
			err := manager.unparkVehicle(parts)
			if err != nil {
				fmt.Println(err)
			}
		case "display":
			err := manager.display(parts)
			if err != nil {
				fmt.Println(err)
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("invalid command")
		}
	}
}

func (m *ParkingLotManager) create(parts []string) error {
	if len(parts) != 4 {
		return fmt.Errorf("invalid command")
	}
	parkingLotID := parts[1]
	noOfFloors := parts[2]
	noOfSlotsPerFloor := parts[3]
	floors, err1 := strconv.Atoi(noOfFloors)
	slots, err2 := strconv.Atoi(noOfSlotsPerFloor)
	if err1 != nil || err2 != nil {
		return fmt.Errorf("invalid command")
	}
	m.parkingLots[parkingLotID] = parkinglot.NewParkingLot(parkingLotID, floors, slots)
	m.primaryId = parkingLotID
	fmt.Printf("Command: create_parking_lot\nparking_lot_id: %s\nno_of_floors: %d\nno_of_slots_per_floor: %d\n",
		parkingLotID, floors, slots)
	return nil
}

func (m *ParkingLotManager) display(parts []string) error {
	if len(parts) != 3 {
		return fmt.Errorf("invalid command")
	}
	displayType := parts[1]
	vehicleType := parts[2]
	vt, err := vehicle.ParseVehicleType(vehicleType)
	if err != nil {
		return fmt.Errorf("invalid vehicle type")
	}
	lot := m.parkingLots[m.primaryId]

	switch displayType {
	case "free_count":
		lot.DisplayFreeCount(vt)
	case "free_slots":
		lot.DisplayFreeSlots(vt)
	case "occupied_slots":
		lot.DisplayOccupiedSlots(vt)
	default:
		return fmt.Errorf("invalid command")

	}
	return nil
}

func (m *ParkingLotManager) parkVehicle(parts []string) error {
	if len(parts) != 4 {
		return fmt.Errorf("invalid command")
	}
	vehicleType := parts[1]
	regNo := parts[2]
	color := parts[3]
	vt, err := vehicle.ParseVehicleType(vehicleType)
	if err != nil {
		return fmt.Errorf("invalid vehicle type")
	}
	lot := m.parkingLots[m.primaryId]
	v := vehicle.NewVehicle(vt, regNo, color)
	ticketId, err := lot.ParkVehicle(v)
	if err != nil {
		return err
	}
	fmt.Printf("Parked vehicle. Ticket ID: %s\n", ticketId)
	return nil
}

func (m *ParkingLotManager) unparkVehicle(parts []string) error {
	if len(parts) != 2 {
		return fmt.Errorf("invalid command")
	}
	ticketID := parts[1]
	// Ticket format: PR<parking_lot_id>_<floor_no>_<slot_no>
	ticketParts := strings.Split(ticketID, "_")
	if len(ticketParts) != 3 {
		return fmt.Errorf("invalid ticket id")
	}
	parkingLotID := ticketParts[0]
	floorNo, err1 := strconv.Atoi(ticketParts[1])
	slotNo := ticketParts[2]
	if err1 != nil {
		return fmt.Errorf("invalid ticket id")
	}
	fmt.Println(parkingLotID, floorNo, slotNo)
	fmt.Println(m.parkingLots)
	lot, ok := m.parkingLots[parkingLotID]
	if !ok {
		return fmt.Errorf("parking lot not found")
	}
	var found bool
	for _, floor := range lot.ParkingFloors {
		if floor.FloorNumber == floorNo {
			for _, slot := range floor.ParkingSlots {
				if slot.ID == slotNo && slot.Status == parkingslot.StatusOccupied {
					slot.Status = parkingslot.StatusAvailable
					slot.Vehicle = nil
					found = true
					break
				}
			}
			break
		}
	}
	if found {
		fmt.Printf("Unparked vehicle with Ticket ID: %s\n", ticketID)
		return nil
	}
	return fmt.Errorf("slot not found or not occupied")
}
