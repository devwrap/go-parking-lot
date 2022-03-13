package model

import "fmt"

type ParkingLot struct {
	size         int
	parkingSlots []ParkingSlot
}

func (pl *ParkingLot) CreateParkingLot(size int) {
	pl.size = size
	pl.parkingSlots = make([]ParkingSlot, size)
}

func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) {
	spot := pl.getRecentSpot()
	if spot == -1 {
		fmt.Println("Parking lot is full!!")
		return
	} else {
		pl.parkingSlots[spot].vehicle = vehicle
		pl.parkingSlots[spot].available = false
	}
}

func (pl *ParkingLot) UnParkVehicle() {

}

func (pl ParkingLot) getRecentSpot() int {
	for i := 0; i < pl.size; i++ {
		if pl.parkingSlots[i].available {
			return i
		}
	}
	return -1
}
