package main

import "fmt"

type RailroadWidthChecker interface {
	CheckRailWidth() int
}

type Railroad struct {
	Width int
}

func (r *Railroad) IsCorrectSizeTrain(p RailroadWidthChecker) bool {
	return p.CheckRailWidth() != r.Width
}

type Train struct {
	Width int
}

func (t *Train) CheckRailWidth() int {
	return t.Width
}

func main() {
	railroad := Railroad{Width: 10}
	passengerTrain := Train{Width: 10}
	cargoTrain := Train{Width: 15}
	canPassengerTrainGo := railroad.IsCorrectSizeTrain(&passengerTrain)
	fmt.Printf("can passneger train go? %b\n", canPassengerTrainGo)
	canCargoTrainGo := railroad.IsCorrectSizeTrain(&cargoTrain)
	fmt.Printf("can cargo train go? %b\n", canCargoTrainGo)
}
