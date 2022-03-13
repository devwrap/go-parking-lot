package model

import (
	"errors"
	"fmt"
)

type ParkingLot struct {
	size         int
	parkingSlots []ParkingSlot
}

func (pl *ParkingLot) CreateParkingLot(size int) {
	pl.size = size
	pl.parkingSlots = make([]ParkingSlot, size)
}

func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) {
	spot, err := pl.getRecentSpot()
	if err != nil {
		fmt.Println("Parking failed: ", err)
	} else {
		spot.vehicle = vehicle
		spot.available = false
	}
}

func (pl *ParkingLot) UnParkVehicle(vehicleNumber string) {
	spot, err := pl.getParkingSlot(vehicleNumber)
	if err != nil {
		fmt.Println("UnPark failed: ", err)
	} else {
		spot.available = true
		return
	}
}

func (pl ParkingLot) getRecentSpot() (*ParkingSlot, error) {
	for _, parkingSlot := range pl.parkingSlots {
		if parkingSlot.available {
			return &parkingSlot, nil
		}
	}
	for i := 0; i < pl.size; i++ {
	}
	return nil, errors.New("parking lot is full")
}

func (pl *ParkingLot) getParkingSlot(vehicleNumber string) (*ParkingSlot, error) {
	for _, parkingSlot := range pl.parkingSlots {
		if parkingSlot.available == false && parkingSlot.vehicle.number == vehicleNumber {
			return &parkingSlot, nil
		}
	}
	return nil, errors.New("no such car parked")
}
