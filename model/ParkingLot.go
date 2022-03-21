package model

import (
	"errors"
	"fmt"
)

type ParkingLot struct {
	Size         int
	ParkingSlots []ParkingSlot
}

func (pl *ParkingLot) CreateParkingLot(size int) {
	pl.Size = size
	pl.ParkingSlots = make([]ParkingSlot, size)
	for spot, _ := range pl.ParkingSlots {
		pl.ParkingSlots[spot].available = true
	}
}

func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) {
	spotNumber, err := pl.getRecentSpot()
	if err != nil {
		fmt.Println("Parking failed: ", err)
	} else {
		pl.ParkingSlots[spotNumber].vehicle = vehicle
		pl.ParkingSlots[spotNumber].available = false
		fmt.Println("Parked car successfully")
	}
}

func (pl *ParkingLot) UnParkVehicle(vehicleNumber string) {
	spotNumber, err := pl.getParkingSlot(vehicleNumber)
	if err != nil {
		fmt.Println("UnPark failed: ", err)
	} else {
		pl.ParkingSlots[spotNumber].available = true
		fmt.Println("Parked car successfully")
	}
}

func (pl ParkingLot) getRecentSpot() (int, error) {
	for spotNumber, parkingSlot := range pl.ParkingSlots {
		if parkingSlot.available {
			return spotNumber, nil
		}
	}
	for i := 0; i < pl.Size; i++ {
	}
	return -1, errors.New("parking lot is full")
}

func (pl *ParkingLot) getParkingSlot(vehicleNumber string) (int, error) {
	for spotNumber, parkingSlot := range pl.ParkingSlots {
		if parkingSlot.available == false && parkingSlot.vehicle.Number == vehicleNumber {
			return spotNumber, nil
		}
	}
	return -1, errors.New("no such car parked")
}
