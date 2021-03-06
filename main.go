package main

import (
	"bufio"
	"fmt"
	m "go-parking-lot/model"
	"log"
	"os"
	"strconv"
	"strings"
)

var app m.ParkingLot

func main() {
	startCLI()
}

func startCLI() {
	fmt.Println("Start by creating parking lot..")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		action, value := getCommand(scanner.Text())
		performAction(action, value)
	}
}

func startApp() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		action, value := getCommand(scanner.Text())
		performAction(action, value)
	}
}

func getCommand(command string) (string, string) {
	commands := strings.SplitN(command, ",", 2)
	return strings.TrimSpace(commands[0]), strings.TrimSpace(commands[1])
}

func performAction(action, value string) {
	switch action {
	case "create":
		{
			// Check if the app is already started
			if app.Size != 0 {
				fmt.Println("App already started!")
				return
			}
			fmt.Println("Starting parking lot of size: " + value)
			size, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			app.CreateParkingLot(size)
		}
	case "park":
		{
			value := strings.SplitN(value, ",", 2)
			vehicle := new(m.Vehicle)
			vehicle.Number = value[0]
			vehicle.Color = value[1]
			app.ParkVehicle(*vehicle)
		}
	case "unpark":
		{
			app.UnParkVehicle(value)
		}
	case "exit":
		{
			os.Exit(1)
		}
	default:
		fmt.Println("no command found")
	}
}
