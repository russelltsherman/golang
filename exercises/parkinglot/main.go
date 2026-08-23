package main

import (
	"errors"
	"fmt"
)

type Spot int
type Vehicle int

const (
	Sm  Spot    = 1
	Rg  Spot    = 2
	Lg  Spot    = 3
	Mc  Vehicle = 1
	Car Vehicle = 2
	Van Vehicle = 3
)

type Parkinglot struct {
	SmSpot       int
	RgSpot       int
	LgSpot       int
	SmSpotFilled int
	RgSpotFilled int
	LgSpotFilled int
}

func (p Parkinglot) isAvailable(spot Spot, numSpot int) bool {
	switch spot {
	case Sm:
		return (p.SmSpotFilled + numSpot) <= p.SmSpot
	case Rg:
		return (p.RgSpotFilled + numSpot) <= p.RgSpot
	case Lg:
		return (p.LgSpotFilled + numSpot) <= p.LgSpot
	default:
		return false
	}
}

// * when the parking lot is empty
func (p Parkinglot) isEmpty() bool {
	return p.SmSpotFilled == 0 && p.RgSpotFilled == 0 && p.LgSpotFilled == 0
}

// * when the parking lot is full
func (p Parkinglot) ifFull() bool {
	return p.SmSpotFilled == p.SmSpot && p.RgSpotFilled == p.RgSpot && p.LgSpotFilled == p.LgSpot
}

// * how many spots vans are taking up
func (p Parkinglot) occupancy(vehicle Vehicle) int {
	return 0
}

// park a type of vehicle
func (p *Parkinglot) park(vehicle Vehicle, spot Spot) error {
	numSpot := 1

	switch spot {
	case Sm:
		if vehicle != Mc {
			return errors.New("vehicle does not fit")
		}
	case Rg:
		if vehicle == Van {
			numSpot = 3
		}
	}

	if p.isAvailable(spot, numSpot) == false {
		return errors.New("Spot not available")
	}

	switch vehicle {
	case Mc:
		if spot == Sm {
			(*p).SmSpotFilled += numSpot
		}
		if spot == Rg {
			(*p).RgSpotFilled += numSpot
		}
		if spot == Lg {
			(*p).LgSpotFilled += numSpot
		}
	case Car:
		if spot == Rg {
			(*p).RgSpotFilled += numSpot
		}
		if spot == Lg {
			(*p).LgSpotFilled += numSpot
		}
	case Van:
		if spot == Rg {
			(*p).RgSpotFilled += numSpot
		}
		if spot == Lg {
			(*p).LgSpotFilled += numSpot
		}

	default:
		return errors.New("can't park that vehicle type")
	}

	return nil
}

// * how many spots are remaining
func (p Parkinglot) spotsRemaining() int {
	return (p.SmSpot - p.SmSpotFilled) + (p.RgSpot - p.RgSpotFilled) + (p.LgSpot - p.LgSpotFilled)
}

// * how many total spots are in the parking lot
func (p Parkinglot) spotsTotal() int {
	return p.SmSpot + p.RgSpot + p.LgSpot
}

// // vacate a type of vehicle from a spot
// func (p Parkinglot) vacate(vehicle Vehicle, spot Spot) error {
// 	return nil
// }

func main() {
	// fmt.Print("Hello")
	p := Parkinglot{
		SmSpot: 15,
		RgSpot: 100,
		LgSpot: 6,
	}

	fmt.Println(p)

	err := p.park(Mc, Sm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	err = p.park(Mc, Rg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	err = p.park(Mc, Lg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	err = p.park(Van, Lg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	err = p.park(Van, Rg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	err = p.park(Van, Sm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

}
