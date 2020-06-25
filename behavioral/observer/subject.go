package observer

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}
