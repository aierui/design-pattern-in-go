package observer

type observer interface {
	update(string)
	getID() string
}
