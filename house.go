package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type House struct {
	NumberOfRooms int
	City, Address string
	Price         float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var houses []House
	input := ""

	for input != "done" {
		fmt.Println("Number of rooms, City, Address, Price")
		fmt.Println("Type 'finished' when done")

		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "finished" {
			break
		}

		parts := strings.Split(input, ",")
		if len(parts) != 4 {
			fmt.Println("Invalid input, try again.")
			continue
		}

		numberOfRooms, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		city := strings.TrimSpace(parts[1])
		address := strings.TrimSpace(parts[2])
		price, err2 := strconv.ParseFloat(strings.TrimSpace(parts[3]), 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Error in number conversion, try again.")
			continue
		}

		houses = append(houses, House{
			NumberOfRooms: numberOfRooms,
			City:          city,
			Address:       address,
			Price:         price,
		})
	}

	for _, house := range houses {
		fmt.Printf("%s\t %s\t %d Rooms\t $%.2f\t\n", house.Address, house.City, house.NumberOfRooms, house.Price)
	}
}
