package responsibility

type department interface {
	execute(*patient)
	setNext(department)
}
