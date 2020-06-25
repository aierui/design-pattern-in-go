package mediator

type train interface {
	requestArrival()
	departure()
	permitArrival()
}
