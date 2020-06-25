package mediator

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (g *passengerTrain) departure() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyFree()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}
