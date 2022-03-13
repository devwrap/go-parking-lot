package main

import (
	"fmt"
	m "parking-lot/model"
)

func main() {
	fmt.Println("Starting parking lot...")

	parkingLot := new(m.ParkingLot)
	parkingLot.CreateParkingLot(10)

}
