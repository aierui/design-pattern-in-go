package mediator

type mediator interface {
	canLand(train) bool
	notifyFree()
}
